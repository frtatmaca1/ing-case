/*
* We need a documentation about redis and cache mechanism.
* We need meaningful validations about parameters.
* We need to add success, warning, error logs.
* We need to add Try/Catch/Callback/Retry mechanism.
* We need to set Redis client authentication.
* We need to listen get/set/del/flush func response(callback, try/catch)
* We need to develop func as async/await/promise
* I there are timeout settings, we need to set them.
* We need to put the healtcheck (readinessa) redis, not getting the request before redis is ready, or setting the default of the place calling the redis
* We need to collect metrics for monitoring(to endpoint for start date, end date)(Wrapping functions with tracing)
 */

var redisService = require('lib/redisService');

/*
   * it is need to be "let". instead of const.
   * Or if we need immutable object, we don't need init
*/
const cache;


/*
   * Probably redis init func have a timeout time, we need to add it.
   * Redis init func can be return error, we can encounter errors. we need to handle them and return the exception message to client.
   * we need to call the external services with async, await or promise.
*/
function init(opts) {
    cache = redisService.init({
        port: opts.port,
        host: opts.host
    });
}

/*
   * this func can return error, we need to handle them
   * before to call get func, we can validate key, it can be null or undefined
   * we need to call the external services with async, await or promise.
*/
function get(key) {
    return cache.get(key);
}

/*
   * this func can return error, we need to handle them
   * before to call get func, we can validate key, value, ttl, it can be null or undefined
   * we need to call the external services with async, await or promise.
 */
function set(key, value, ttl) {
    return cache.set(key, value, ttl || opts.stdTTL);
}

/*
   * this func can return error, we need to handle them
   * before to call get func, we can validate key, it can be null or undefined
   * we need to call the external services with async, await or promise.
*/
function del(key) {
    return cache.del(key);
}

/*
   * this func can return error, we need to handle them
   * we need to call the external services with async, await or promise.
*/
function flush() {
    return cache.flush();
}

function getCache(opts) {
    return {
        get: get,
        set: set,
        del: del,
        flush: flush
    };
}

module.exports = {
    init: init,
    getCache: getCache
};