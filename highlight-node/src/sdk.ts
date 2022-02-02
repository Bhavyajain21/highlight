import { Highlight } from '.';
import { NodeOptions } from './types';

export const HIGHLIGHT_REQUEST_HEADER = 'x-highlight-request';

export interface HighlightInterface {
    init: (options: NodeOptions) => void;
    isInitialized: () => boolean;
    consumeError: (
        error: Error,
        secureSessionId: string,
        requestId: string
    ) => void;
}

var highlight_obj: Highlight;
export const H: HighlightInterface = {
    init: (options: NodeOptions = {}) => {
        try {
            highlight_obj = new Highlight(options);
        } catch (e) {
            console.log('highlight-next init error: ', e);
        }
    },
    isInitialized: () => (highlight_obj ? true : false),
    consumeError: (
        error: Error,
        secureSessionId: string,
        requestId: string
    ) => {
        try {
            highlight_obj.consumeCustomError(error, secureSessionId, requestId);
        } catch (e) {
            console.log('highlight-next error: ', e);
        }
    },
};
