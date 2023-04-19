class PromiseState {
    static PENDING = 0;
    static RESOLVED = 1;
    static REJECTED = 2;
}

export class MyPromise {
    state = PromiseState.PENDING;

    callbacks = [];

    errorHandler;

    constructor(callback) {
        callback(this.resolve, this.reject);
    }

    // Public methods must be bound with `this` to avoid a `lost this` problem.
    resolve = function (data) {
        this.state = PromiseState.RESOLVED;
        this.#runCallbacks(data);
    }.bind(this);

    reject = function (errMsg) {
        this.state = PromiseState.REJECTED;
        this.errorHandler(errMsg);
    }.bind(this);

    then = function (callback) {
        this.callbacks.push(callback);
        return this;
    }.bind(this);

    catch = function (handler) {
        this.errorHandler = handler;
    }.bind(this);

    #runCallbacks(data) {
        this.callbacks.reduce((prev, next) => next(prev), data);
    }
}
