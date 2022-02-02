import { NextApiHandler, NextApiRequest, NextApiResponse } from 'next';
import { NodeOptions, H, HIGHLIGHT_REQUEST_HEADER } from '@highlight-run/node';

// This is the same as the `NextApiHandler` type, except instead of having a return type of `void | Promise<void>`, it's
// only `Promise<void>`, because wrapped handlers are always async
export type WrappedNextApiHandler = (
    req: NextApiRequest,
    res: NextApiResponse
) => Promise<void>;

export type AugmentedNextApiResponse = NextApiResponse;

export const Highlight =
    (options: NodeOptions = {}) =>
    (origHandler: NextApiHandler): WrappedNextApiHandler => {
        return async (req, res) => {
            try {
                return await origHandler(req, res);
            } catch (e) {
                if (req.headers && req.headers[HIGHLIGHT_REQUEST_HEADER]) {
                    const [secureSessionId, requestId] =
                        `${req.headers[HIGHLIGHT_REQUEST_HEADER]}`.split('/');
                    if (secureSessionId && requestId) {
                        if (e instanceof Error) {
                            if (!H.isInitialized()) {
                                H.init(options);
                            }
                            H.consumeError(e, secureSessionId, requestId);
                        }
                    }
                }

                // Because we're going to finish and send the transaction before passing the error onto nextjs, it won't yet
                // have had a chance to set the status to 500, so unless we do it ourselves now, we'll incorrectly report that
                // the transaction was error-free
                res.statusCode = 500;
                res.statusMessage = 'Internal Server Error';

                // We rethrow here so that nextjs can do with the error whatever it would normally do. (Sometimes "whatever it
                // would normally do" is to allow the error to bubble up to the global handlers - another reason we need to mark
                // the error as already having been captured.)
                throw e;
            }
        };
    };

// this code was assisted by https://github.com/getsentry/sentry-javascript/blob/2f2d099dcc668f12109913caf36097f73a1bc966/packages/nextjs/src/utils/withSentry.ts
