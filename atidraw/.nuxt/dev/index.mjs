import process from 'node:process';globalThis._importMeta_={url:import.meta.url,env:process.env};import { tmpdir } from 'node:os';
import { defineEventHandler, handleCacheHeaders, splitCookiesString, createEvent, fetchWithEvent, isEvent, eventHandler, setHeaders, sendRedirect, proxyRequest, getRequestHeader, setResponseHeaders, setResponseStatus, send, getRequestHeaders, setResponseHeader, appendResponseHeader, getRequestURL, getResponseHeader, removeResponseHeader, createError, getQuery as getQuery$1, readBody, getResponseStatus, getCookie, deleteCookie, setCookie, useSession, createApp, createRouter as createRouter$1, toNodeListener, lazyEventHandler, getRouterParam, readFormData, setHeader, getRequestFingerprint, getResponseStatusText } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/h3@1.15.4/node_modules/h3/dist/index.mjs';
import { Server } from 'node:http';
import { resolve, dirname, join } from 'node:path';
import crypto$1 from 'node:crypto';
import { parentPort, threadId } from 'node:worker_threads';
import { escapeHtml } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/@vue+shared@3.5.26/node_modules/@vue/shared/dist/shared.cjs.js';
import { blob, ensureBlob } from 'file:///Users/mojito/Code/atidraw/.nuxt/hub/blob.mjs';
import { generateText } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/ai@5.0.115_zod@4.2.1/node_modules/ai/dist/index.mjs';
import { parseURL, withoutBase, joinURL, getQuery, withQuery, withTrailingSlash, decodePath, withLeadingSlash, withoutTrailingSlash, joinRelativeURL } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/ufo@1.6.1/node_modules/ufo/dist/index.mjs';
import defu, { defuFn, defu as defu$1 } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/defu@6.1.4/node_modules/defu/dist/defu.mjs';
import { FetchError, createFetch, Headers as Headers$1 } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/ofetch@1.5.1/node_modules/ofetch/dist/node.mjs';
import { snakeCase, upperFirst } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/scule@1.3.0/node_modules/scule/dist/index.mjs';
import { getRandomValues } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/uncrypto@0.1.3/node_modules/uncrypto/dist/crypto.node.mjs';
import { createRenderer, getRequestDependencies, getPreloadLinks, getPrefetchLinks } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/vue-bundle-renderer@2.2.0/node_modules/vue-bundle-renderer/dist/runtime.mjs';
import destr, { destr as destr$1 } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/destr@2.0.5/node_modules/destr/dist/index.mjs';
import { createHooks } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/hookable@5.5.3/node_modules/hookable/dist/index.mjs';
import { fetchNodeRequestHandler, callNodeRequestHandler } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/node-mock-http@1.0.3/node_modules/node-mock-http/dist/index.mjs';
import { createStorage, prefixStorage } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/unstorage@1.17.3_@netlify+blobs@9.1.2_@vercel+blob@2.0.0_aws4fetch@1.0.20_db0@0.3.4_ioredis@5.8.2/node_modules/unstorage/dist/index.mjs';
import unstorage_47drivers_47fs from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/unstorage@1.17.3_@netlify+blobs@9.1.2_@vercel+blob@2.0.0_aws4fetch@1.0.20_db0@0.3.4_ioredis@5.8.2/node_modules/unstorage/drivers/fs.mjs';
import { digest, hash as hash$1 } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/ohash@2.0.11/node_modules/ohash/dist/index.mjs';
import { klona } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/klona@2.0.6/node_modules/klona/dist/index.mjs';
import { AsyncLocalStorage } from 'node:async_hooks';
import { getContext } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/unctx@2.4.1/node_modules/unctx/dist/index.mjs';
import { toRouteMatcher, createRouter } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/radix3@1.1.2/node_modules/radix3/dist/index.mjs';
import { readFile } from 'node:fs/promises';
import consola, { consola as consola$1 } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/consola@3.4.2/node_modules/consola/dist/index.mjs';
import { ErrorParser } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/youch-core@0.3.3/node_modules/youch-core/build/index.js';
import { Youch } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/youch@4.1.0-beta.11/node_modules/youch/build/index.js';
import { SourceMapConsumer } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/source-map@0.7.6/node_modules/source-map/source-map.js';
import { stringify, uneval } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/devalue@5.6.1/node_modules/devalue/index.js';
import { captureRawStackTrace, parseRawStackTrace } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/errx@0.1.0/node_modules/errx/dist/index.js';
import { isVNode, toValue, isRef } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/vue@3.5.26_typescript@5.8.3/node_modules/vue/index.mjs';
import { promises } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { dirname as dirname$1, resolve as resolve$1, basename } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/pathe@2.0.3/node_modules/pathe/dist/index.mjs';
import { createHead as createHead$1, propsToString, renderSSRHead } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/unhead@2.0.19/node_modules/unhead/dist/server.mjs';
import process$1 from 'node:process';
import { renderToString } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/vue@3.5.26_typescript@5.8.3/node_modules/vue/server-renderer/index.mjs';
import { walkResolver } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/unhead@2.0.19/node_modules/unhead/dist/utils.mjs';
import { getIcons } from 'file:///Users/mojito/Code/atidraw/node_modules/.pnpm/@iconify+utils@3.1.0/node_modules/@iconify/utils/lib/index.js';
import { collections } from 'file:///Users/mojito/Code/atidraw/.nuxt/nuxt-icon-server-bundle.mjs';

const serverAssets = [{"baseName":"server","dir":"/Users/mojito/Code/atidraw/server/assets"}];

const assets$1 = createStorage();

for (const asset of serverAssets) {
  assets$1.mount(asset.baseName, unstorage_47drivers_47fs({ base: asset.dir, ignore: (asset?.ignore || []) }));
}

const storage = createStorage({});

storage.mount('/assets', assets$1);

storage.mount('root', unstorage_47drivers_47fs({"driver":"fs","readOnly":true,"base":"/Users/mojito/Code/atidraw","watchOptions":{"ignored":[null]}}));
storage.mount('src', unstorage_47drivers_47fs({"driver":"fs","readOnly":true,"base":"/Users/mojito/Code/atidraw/server","watchOptions":{"ignored":[null]}}));
storage.mount('build', unstorage_47drivers_47fs({"driver":"fs","readOnly":false,"base":"/Users/mojito/Code/atidraw/.nuxt"}));
storage.mount('cache', unstorage_47drivers_47fs({"driver":"fs","readOnly":false,"base":"/Users/mojito/Code/atidraw/.nuxt/cache"}));
storage.mount('data', unstorage_47drivers_47fs({"driver":"fs","base":"/Users/mojito/Code/atidraw/.data/kv"}));

function useStorage(base = "") {
  return base ? prefixStorage(storage, base) : storage;
}

const Hasher = /* @__PURE__ */ (() => {
  class Hasher2 {
    buff = "";
    #context = /* @__PURE__ */ new Map();
    write(str) {
      this.buff += str;
    }
    dispatch(value) {
      const type = value === null ? "null" : typeof value;
      return this[type](value);
    }
    object(object) {
      if (object && typeof object.toJSON === "function") {
        return this.object(object.toJSON());
      }
      const objString = Object.prototype.toString.call(object);
      let objType = "";
      const objectLength = objString.length;
      objType = objectLength < 10 ? "unknown:[" + objString + "]" : objString.slice(8, objectLength - 1);
      objType = objType.toLowerCase();
      let objectNumber = null;
      if ((objectNumber = this.#context.get(object)) === void 0) {
        this.#context.set(object, this.#context.size);
      } else {
        return this.dispatch("[CIRCULAR:" + objectNumber + "]");
      }
      if (typeof Buffer !== "undefined" && Buffer.isBuffer && Buffer.isBuffer(object)) {
        this.write("buffer:");
        return this.write(object.toString("utf8"));
      }
      if (objType !== "object" && objType !== "function" && objType !== "asyncfunction") {
        if (this[objType]) {
          this[objType](object);
        } else {
          this.unknown(object, objType);
        }
      } else {
        const keys = Object.keys(object).sort();
        const extraKeys = [];
        this.write("object:" + (keys.length + extraKeys.length) + ":");
        const dispatchForKey = (key) => {
          this.dispatch(key);
          this.write(":");
          this.dispatch(object[key]);
          this.write(",");
        };
        for (const key of keys) {
          dispatchForKey(key);
        }
        for (const key of extraKeys) {
          dispatchForKey(key);
        }
      }
    }
    array(arr, unordered) {
      unordered = unordered === void 0 ? false : unordered;
      this.write("array:" + arr.length + ":");
      if (!unordered || arr.length <= 1) {
        for (const entry of arr) {
          this.dispatch(entry);
        }
        return;
      }
      const contextAdditions = /* @__PURE__ */ new Map();
      const entries = arr.map((entry) => {
        const hasher = new Hasher2();
        hasher.dispatch(entry);
        for (const [key, value] of hasher.#context) {
          contextAdditions.set(key, value);
        }
        return hasher.toString();
      });
      this.#context = contextAdditions;
      entries.sort();
      return this.array(entries, false);
    }
    date(date) {
      return this.write("date:" + date.toJSON());
    }
    symbol(sym) {
      return this.write("symbol:" + sym.toString());
    }
    unknown(value, type) {
      this.write(type);
      if (!value) {
        return;
      }
      this.write(":");
      if (value && typeof value.entries === "function") {
        return this.array(
          [...value.entries()],
          true
          /* ordered */
        );
      }
    }
    error(err) {
      return this.write("error:" + err.toString());
    }
    boolean(bool) {
      return this.write("bool:" + bool);
    }
    string(string) {
      this.write("string:" + string.length + ":");
      this.write(string);
    }
    function(fn) {
      this.write("fn:");
      if (isNativeFunction(fn)) {
        this.dispatch("[native]");
      } else {
        this.dispatch(fn.toString());
      }
    }
    number(number) {
      return this.write("number:" + number);
    }
    null() {
      return this.write("Null");
    }
    undefined() {
      return this.write("Undefined");
    }
    regexp(regex) {
      return this.write("regex:" + regex.toString());
    }
    arraybuffer(arr) {
      this.write("arraybuffer:");
      return this.dispatch(new Uint8Array(arr));
    }
    url(url) {
      return this.write("url:" + url.toString());
    }
    map(map) {
      this.write("map:");
      const arr = [...map];
      return this.array(arr, false);
    }
    set(set) {
      this.write("set:");
      const arr = [...set];
      return this.array(arr, false);
    }
    bigint(number) {
      return this.write("bigint:" + number.toString());
    }
  }
  for (const type of [
    "uint8array",
    "uint8clampedarray",
    "unt8array",
    "uint16array",
    "unt16array",
    "uint32array",
    "unt32array",
    "float32array",
    "float64array"
  ]) {
    Hasher2.prototype[type] = function(arr) {
      this.write(type + ":");
      return this.array([...arr], false);
    };
  }
  function isNativeFunction(f) {
    if (typeof f !== "function") {
      return false;
    }
    return Function.prototype.toString.call(f).slice(
      -15
      /* "[native code] }".length */
    ) === "[native code] }";
  }
  return Hasher2;
})();
function serialize(object) {
  const hasher = new Hasher();
  hasher.dispatch(object);
  return hasher.buff;
}
function hash(value) {
  return digest(typeof value === "string" ? value : serialize(value)).replace(/[-_]/g, "").slice(0, 10);
}

function defaultCacheOptions() {
  return {
    name: "_",
    base: "/cache",
    swr: true,
    maxAge: 1
  };
}
function defineCachedFunction(fn, opts = {}) {
  opts = { ...defaultCacheOptions(), ...opts };
  const pending = {};
  const group = opts.group || "nitro/functions";
  const name = opts.name || fn.name || "_";
  const integrity = opts.integrity || hash([fn, opts]);
  const validate = opts.validate || ((entry) => entry.value !== void 0);
  async function get(key, resolver, shouldInvalidateCache, event) {
    const cacheKey = [opts.base, group, name, key + ".json"].filter(Boolean).join(":").replace(/:\/$/, ":index");
    let entry = await useStorage().getItem(cacheKey).catch((error) => {
      console.error(`[cache] Cache read error.`, error);
      useNitroApp().captureError(error, { event, tags: ["cache"] });
    }) || {};
    if (typeof entry !== "object") {
      entry = {};
      const error = new Error("Malformed data read from cache.");
      console.error("[cache]", error);
      useNitroApp().captureError(error, { event, tags: ["cache"] });
    }
    const ttl = (opts.maxAge ?? 0) * 1e3;
    if (ttl) {
      entry.expires = Date.now() + ttl;
    }
    const expired = shouldInvalidateCache || entry.integrity !== integrity || ttl && Date.now() - (entry.mtime || 0) > ttl || validate(entry) === false;
    const _resolve = async () => {
      const isPending = pending[key];
      if (!isPending) {
        if (entry.value !== void 0 && (opts.staleMaxAge || 0) >= 0 && opts.swr === false) {
          entry.value = void 0;
          entry.integrity = void 0;
          entry.mtime = void 0;
          entry.expires = void 0;
        }
        pending[key] = Promise.resolve(resolver());
      }
      try {
        entry.value = await pending[key];
      } catch (error) {
        if (!isPending) {
          delete pending[key];
        }
        throw error;
      }
      if (!isPending) {
        entry.mtime = Date.now();
        entry.integrity = integrity;
        delete pending[key];
        if (validate(entry) !== false) {
          let setOpts;
          if (opts.maxAge && !opts.swr) {
            setOpts = { ttl: opts.maxAge };
          }
          const promise = useStorage().setItem(cacheKey, entry, setOpts).catch((error) => {
            console.error(`[cache] Cache write error.`, error);
            useNitroApp().captureError(error, { event, tags: ["cache"] });
          });
          if (event?.waitUntil) {
            event.waitUntil(promise);
          }
        }
      }
    };
    const _resolvePromise = expired ? _resolve() : Promise.resolve();
    if (entry.value === void 0) {
      await _resolvePromise;
    } else if (expired && event && event.waitUntil) {
      event.waitUntil(_resolvePromise);
    }
    if (opts.swr && validate(entry) !== false) {
      _resolvePromise.catch((error) => {
        console.error(`[cache] SWR handler error.`, error);
        useNitroApp().captureError(error, { event, tags: ["cache"] });
      });
      return entry;
    }
    return _resolvePromise.then(() => entry);
  }
  return async (...args) => {
    const shouldBypassCache = await opts.shouldBypassCache?.(...args);
    if (shouldBypassCache) {
      return fn(...args);
    }
    const key = await (opts.getKey || getKey)(...args);
    const shouldInvalidateCache = await opts.shouldInvalidateCache?.(...args);
    const entry = await get(
      key,
      () => fn(...args),
      shouldInvalidateCache,
      args[0] && isEvent(args[0]) ? args[0] : void 0
    );
    let value = entry.value;
    if (opts.transform) {
      value = await opts.transform(entry, ...args) || value;
    }
    return value;
  };
}
function cachedFunction(fn, opts = {}) {
  return defineCachedFunction(fn, opts);
}
function getKey(...args) {
  return args.length > 0 ? hash(args) : "";
}
function escapeKey(key) {
  return String(key).replace(/\W/g, "");
}
function defineCachedEventHandler(handler, opts = defaultCacheOptions()) {
  const variableHeaderNames = (opts.varies || []).filter(Boolean).map((h) => h.toLowerCase()).sort();
  const _opts = {
    ...opts,
    getKey: async (event) => {
      const customKey = await opts.getKey?.(event);
      if (customKey) {
        return escapeKey(customKey);
      }
      const _path = event.node.req.originalUrl || event.node.req.url || event.path;
      let _pathname;
      try {
        _pathname = escapeKey(decodeURI(parseURL(_path).pathname)).slice(0, 16) || "index";
      } catch {
        _pathname = "-";
      }
      const _hashedPath = `${_pathname}.${hash(_path)}`;
      const _headers = variableHeaderNames.map((header) => [header, event.node.req.headers[header]]).map(([name, value]) => `${escapeKey(name)}.${hash(value)}`);
      return [_hashedPath, ..._headers].join(":");
    },
    validate: (entry) => {
      if (!entry.value) {
        return false;
      }
      if (entry.value.code >= 400) {
        return false;
      }
      if (entry.value.body === void 0) {
        return false;
      }
      if (entry.value.headers.etag === "undefined" || entry.value.headers["last-modified"] === "undefined") {
        return false;
      }
      return true;
    },
    group: opts.group || "nitro/handlers",
    integrity: opts.integrity || hash([handler, opts])
  };
  const _cachedHandler = cachedFunction(
    async (incomingEvent) => {
      const variableHeaders = {};
      for (const header of variableHeaderNames) {
        const value = incomingEvent.node.req.headers[header];
        if (value !== void 0) {
          variableHeaders[header] = value;
        }
      }
      const reqProxy = cloneWithProxy(incomingEvent.node.req, {
        headers: variableHeaders
      });
      const resHeaders = {};
      let _resSendBody;
      const resProxy = cloneWithProxy(incomingEvent.node.res, {
        statusCode: 200,
        writableEnded: false,
        writableFinished: false,
        headersSent: false,
        closed: false,
        getHeader(name) {
          return resHeaders[name];
        },
        setHeader(name, value) {
          resHeaders[name] = value;
          return this;
        },
        getHeaderNames() {
          return Object.keys(resHeaders);
        },
        hasHeader(name) {
          return name in resHeaders;
        },
        removeHeader(name) {
          delete resHeaders[name];
        },
        getHeaders() {
          return resHeaders;
        },
        end(chunk, arg2, arg3) {
          if (typeof chunk === "string") {
            _resSendBody = chunk;
          }
          if (typeof arg2 === "function") {
            arg2();
          }
          if (typeof arg3 === "function") {
            arg3();
          }
          return this;
        },
        write(chunk, arg2, arg3) {
          if (typeof chunk === "string") {
            _resSendBody = chunk;
          }
          if (typeof arg2 === "function") {
            arg2(void 0);
          }
          if (typeof arg3 === "function") {
            arg3();
          }
          return true;
        },
        writeHead(statusCode, headers2) {
          this.statusCode = statusCode;
          if (headers2) {
            if (Array.isArray(headers2) || typeof headers2 === "string") {
              throw new TypeError("Raw headers  is not supported.");
            }
            for (const header in headers2) {
              const value = headers2[header];
              if (value !== void 0) {
                this.setHeader(
                  header,
                  value
                );
              }
            }
          }
          return this;
        }
      });
      const event = createEvent(reqProxy, resProxy);
      event.fetch = (url, fetchOptions) => fetchWithEvent(event, url, fetchOptions, {
        fetch: useNitroApp().localFetch
      });
      event.$fetch = (url, fetchOptions) => fetchWithEvent(event, url, fetchOptions, {
        fetch: globalThis.$fetch
      });
      event.waitUntil = incomingEvent.waitUntil;
      event.context = incomingEvent.context;
      event.context.cache = {
        options: _opts
      };
      const body = await handler(event) || _resSendBody;
      const headers = event.node.res.getHeaders();
      headers.etag = String(
        headers.Etag || headers.etag || `W/"${hash(body)}"`
      );
      headers["last-modified"] = String(
        headers["Last-Modified"] || headers["last-modified"] || (/* @__PURE__ */ new Date()).toUTCString()
      );
      const cacheControl = [];
      if (opts.swr) {
        if (opts.maxAge) {
          cacheControl.push(`s-maxage=${opts.maxAge}`);
        }
        if (opts.staleMaxAge) {
          cacheControl.push(`stale-while-revalidate=${opts.staleMaxAge}`);
        } else {
          cacheControl.push("stale-while-revalidate");
        }
      } else if (opts.maxAge) {
        cacheControl.push(`max-age=${opts.maxAge}`);
      }
      if (cacheControl.length > 0) {
        headers["cache-control"] = cacheControl.join(", ");
      }
      const cacheEntry = {
        code: event.node.res.statusCode,
        headers,
        body
      };
      return cacheEntry;
    },
    _opts
  );
  return defineEventHandler(async (event) => {
    if (opts.headersOnly) {
      if (handleCacheHeaders(event, { maxAge: opts.maxAge })) {
        return;
      }
      return handler(event);
    }
    const response = await _cachedHandler(
      event
    );
    if (event.node.res.headersSent || event.node.res.writableEnded) {
      return response.body;
    }
    if (handleCacheHeaders(event, {
      modifiedTime: new Date(response.headers["last-modified"]),
      etag: response.headers.etag,
      maxAge: opts.maxAge
    })) {
      return;
    }
    event.node.res.statusCode = response.code;
    for (const name in response.headers) {
      const value = response.headers[name];
      if (name === "set-cookie") {
        event.node.res.appendHeader(
          name,
          splitCookiesString(value)
        );
      } else {
        if (value !== void 0) {
          event.node.res.setHeader(name, value);
        }
      }
    }
    return response.body;
  });
}
function cloneWithProxy(obj, overrides) {
  return new Proxy(obj, {
    get(target, property, receiver) {
      if (property in overrides) {
        return overrides[property];
      }
      return Reflect.get(target, property, receiver);
    },
    set(target, property, value, receiver) {
      if (property in overrides) {
        overrides[property] = value;
        return true;
      }
      return Reflect.set(target, property, value, receiver);
    }
  });
}
const cachedEventHandler = defineCachedEventHandler;

const defineAppConfig = (config) => config;

const appConfig0 = defineAppConfig({
  ui: {
    colors: {
      primary: "blue",
      gray: "neutral"
    }
  }
});

const inlineAppConfig = {
  "nuxt": {},
  "ui": {
    "colors": {
      "primary": "green",
      "secondary": "blue",
      "success": "green",
      "info": "blue",
      "warning": "yellow",
      "error": "red",
      "neutral": "slate"
    },
    "icons": {
      "arrowDown": "i-lucide-arrow-down",
      "arrowLeft": "i-lucide-arrow-left",
      "arrowRight": "i-lucide-arrow-right",
      "arrowUp": "i-lucide-arrow-up",
      "caution": "i-lucide-circle-alert",
      "check": "i-lucide-check",
      "chevronDoubleLeft": "i-lucide-chevrons-left",
      "chevronDoubleRight": "i-lucide-chevrons-right",
      "chevronDown": "i-lucide-chevron-down",
      "chevronLeft": "i-lucide-chevron-left",
      "chevronRight": "i-lucide-chevron-right",
      "chevronUp": "i-lucide-chevron-up",
      "close": "i-lucide-x",
      "copy": "i-lucide-copy",
      "copyCheck": "i-lucide-copy-check",
      "dark": "i-lucide-moon",
      "drag": "i-lucide-grip-vertical",
      "ellipsis": "i-lucide-ellipsis",
      "error": "i-lucide-circle-x",
      "external": "i-lucide-arrow-up-right",
      "eye": "i-lucide-eye",
      "eyeOff": "i-lucide-eye-off",
      "file": "i-lucide-file",
      "folder": "i-lucide-folder",
      "folderOpen": "i-lucide-folder-open",
      "hash": "i-lucide-hash",
      "info": "i-lucide-info",
      "light": "i-lucide-sun",
      "loading": "i-lucide-loader-circle",
      "menu": "i-lucide-menu",
      "minus": "i-lucide-minus",
      "panelClose": "i-lucide-panel-left-close",
      "panelOpen": "i-lucide-panel-left-open",
      "plus": "i-lucide-plus",
      "reload": "i-lucide-rotate-ccw",
      "search": "i-lucide-search",
      "stop": "i-lucide-square",
      "success": "i-lucide-circle-check",
      "system": "i-lucide-monitor",
      "tip": "i-lucide-lightbulb",
      "upload": "i-lucide-upload",
      "warning": "i-lucide-triangle-alert"
    },
    "tv": {
      "twMergeConfig": {}
    }
  },
  "icon": {
    "provider": "server",
    "class": "",
    "aliases": {},
    "iconifyApiEndpoint": "https://api.iconify.design",
    "localApiEndpoint": "/api/_nuxt_icon",
    "fallbackToApi": true,
    "cssSelectorPrefix": "i-",
    "cssWherePseudo": true,
    "cssLayer": "components",
    "mode": "css",
    "attrs": {
      "aria-hidden": true
    },
    "collections": [
      "academicons",
      "akar-icons",
      "ant-design",
      "arcticons",
      "basil",
      "bi",
      "bitcoin-icons",
      "bpmn",
      "brandico",
      "bx",
      "bxl",
      "bxs",
      "bytesize",
      "carbon",
      "catppuccin",
      "cbi",
      "charm",
      "ci",
      "cib",
      "cif",
      "cil",
      "circle-flags",
      "circum",
      "clarity",
      "codicon",
      "covid",
      "cryptocurrency",
      "cryptocurrency-color",
      "dashicons",
      "devicon",
      "devicon-plain",
      "ei",
      "el",
      "emojione",
      "emojione-monotone",
      "emojione-v1",
      "entypo",
      "entypo-social",
      "eos-icons",
      "ep",
      "et",
      "eva",
      "f7",
      "fa",
      "fa-brands",
      "fa-regular",
      "fa-solid",
      "fa6-brands",
      "fa6-regular",
      "fa6-solid",
      "fad",
      "fe",
      "feather",
      "file-icons",
      "flag",
      "flagpack",
      "flat-color-icons",
      "flat-ui",
      "flowbite",
      "fluent",
      "fluent-emoji",
      "fluent-emoji-flat",
      "fluent-emoji-high-contrast",
      "fluent-mdl2",
      "fontelico",
      "fontisto",
      "formkit",
      "foundation",
      "fxemoji",
      "gala",
      "game-icons",
      "geo",
      "gg",
      "gis",
      "gravity-ui",
      "gridicons",
      "grommet-icons",
      "guidance",
      "healthicons",
      "heroicons",
      "heroicons-outline",
      "heroicons-solid",
      "hugeicons",
      "humbleicons",
      "ic",
      "icomoon-free",
      "icon-park",
      "icon-park-outline",
      "icon-park-solid",
      "icon-park-twotone",
      "iconamoon",
      "iconoir",
      "icons8",
      "il",
      "ion",
      "iwwa",
      "jam",
      "la",
      "lets-icons",
      "line-md",
      "logos",
      "ls",
      "lucide",
      "lucide-lab",
      "mage",
      "majesticons",
      "maki",
      "map",
      "marketeq",
      "material-symbols",
      "material-symbols-light",
      "mdi",
      "mdi-light",
      "medical-icon",
      "memory",
      "meteocons",
      "mi",
      "mingcute",
      "mono-icons",
      "mynaui",
      "nimbus",
      "nonicons",
      "noto",
      "noto-v1",
      "octicon",
      "oi",
      "ooui",
      "openmoji",
      "oui",
      "pajamas",
      "pepicons",
      "pepicons-pencil",
      "pepicons-pop",
      "pepicons-print",
      "ph",
      "pixelarticons",
      "prime",
      "ps",
      "quill",
      "radix-icons",
      "raphael",
      "ri",
      "rivet-icons",
      "si-glyph",
      "simple-icons",
      "simple-line-icons",
      "skill-icons",
      "solar",
      "streamline",
      "streamline-emojis",
      "subway",
      "svg-spinners",
      "system-uicons",
      "tabler",
      "tdesign",
      "teenyicons",
      "token",
      "token-branded",
      "topcoat",
      "twemoji",
      "typcn",
      "uil",
      "uim",
      "uis",
      "uit",
      "uiw",
      "unjs",
      "vaadin",
      "vs",
      "vscode-icons",
      "websymbol",
      "weui",
      "whh",
      "wi",
      "wpf",
      "zmdi",
      "zondicons"
    ],
    "fetchTimeout": 1500
  }
};

const appConfig = defuFn(appConfig0, inlineAppConfig);

function getEnv(key, opts) {
  const envKey = snakeCase(key).toUpperCase();
  return destr(
    process.env[opts.prefix + envKey] ?? process.env[opts.altPrefix + envKey]
  );
}
function _isObject(input) {
  return typeof input === "object" && !Array.isArray(input);
}
function applyEnv(obj, opts, parentKey = "") {
  for (const key in obj) {
    const subKey = parentKey ? `${parentKey}_${key}` : key;
    const envValue = getEnv(subKey, opts);
    if (_isObject(obj[key])) {
      if (_isObject(envValue)) {
        obj[key] = { ...obj[key], ...envValue };
        applyEnv(obj[key], opts, subKey);
      } else if (envValue === void 0) {
        applyEnv(obj[key], opts, subKey);
      } else {
        obj[key] = envValue ?? obj[key];
      }
    } else {
      obj[key] = envValue ?? obj[key];
    }
    if (opts.envExpansion && typeof obj[key] === "string") {
      obj[key] = _expandFromEnv(obj[key]);
    }
  }
  return obj;
}
const envExpandRx = /\{\{([^{}]*)\}\}/g;
function _expandFromEnv(value) {
  return value.replace(envExpandRx, (match, key) => {
    return process.env[key] || match;
  });
}

const _inlineRuntimeConfig = {
  "app": {
    "baseURL": "/",
    "buildId": "dev",
    "buildAssetsDir": "/_nuxt/",
    "cdnURL": ""
  },
  "nitro": {
    "envPrefix": "NUXT_",
    "routeRules": {
      "/__nuxt_error": {
        "cache": false
      },
      "/drawings/**": {
        "isr": true
      },
      "/_nuxt/builds/meta/**": {
        "headers": {
          "cache-control": "public, max-age=31536000, immutable"
        }
      },
      "/_nuxt/builds/**": {
        "headers": {
          "cache-control": "public, max-age=1, immutable"
        }
      }
    }
  },
  "public": {
    "apiBase": "http://localhost:18888",
    "ossUrl": "http://124.70.84.161:6900/",
    "hub": {},
    "auth": {
      "loadStrategy": "server-first"
    }
  },
  "hub": {
    "dir": "/Users/mojito/Code/atidraw/.data",
    "hosting": "",
    "blob": {
      "driver": "fs",
      "dir": "/Users/mojito/Code/atidraw/.data/blob"
    },
    "cache": false,
    "db": false,
    "kv": false
  },
  "icon": {
    "serverKnownCssClasses": []
  },
  "session": {
    "name": "nuxt-session",
    "password": "",
    "cookie": {
      "sameSite": "lax"
    }
  },
  "hash": {
    "scrypt": {}
  },
  "webauthn": {
    "register": {},
    "authenticate": {}
  },
  "oauth": {
    "gitea": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": "",
      "baseURL": ""
    },
    "github": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "gitlab": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": "",
      "baseURL": "https://gitlab.com"
    },
    "spotify": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "google": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "twitch": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "auth0": {
      "clientId": "",
      "clientSecret": "",
      "domain": "",
      "audience": "",
      "redirectURL": ""
    },
    "workos": {
      "clientId": "",
      "clientSecret": "",
      "connectionId": "",
      "screenHint": "",
      "redirectURL": ""
    },
    "microsoft": {
      "clientId": "",
      "clientSecret": "",
      "tenant": "",
      "scope": [],
      "authorizationURL": "",
      "tokenURL": "",
      "userURL": "",
      "redirectURL": ""
    },
    "azureb2c": {
      "clientId": "",
      "policy": "",
      "tenant": "",
      "scope": [],
      "authorizationURL": "",
      "tokenURL": "",
      "userURL": "",
      "redirectURL": ""
    },
    "discord": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "battledotnet": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "bluesky": {
      "clientMetadataFilename": "",
      "clientName": "",
      "clientUri": "",
      "logoUri": "",
      "policyUri": "",
      "tosUri": "",
      "scope": [
        "atproto"
      ],
      "grantTypes": [
        "authorization_code"
      ],
      "responseTypes": [
        "code"
      ],
      "applicationType": "web",
      "redirectUris": "",
      "dpopBoundAccessTokens": true,
      "tokenEndpointAuthMethod": "none"
    },
    "keycloak": {
      "clientId": "",
      "clientSecret": "",
      "serverUrl": "",
      "serverUrlInternal": "",
      "realm": "",
      "redirectURL": ""
    },
    "linear": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "linkedin": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "cognito": {
      "clientId": "",
      "clientSecret": "",
      "region": "",
      "userPoolId": "",
      "redirectURL": ""
    },
    "facebook": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "instagram": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "paypal": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "steam": {
      "apiKey": "",
      "redirectURL": ""
    },
    "x": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "xsuaa": {
      "clientId": "",
      "clientSecret": "",
      "domain": "",
      "redirectURL": ""
    },
    "vk": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "yandex": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "tiktok": {
      "clientKey": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "dropbox": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "polar": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "zitadel": {
      "clientId": "",
      "clientSecret": "",
      "domain": "",
      "redirectURL": ""
    },
    "authentik": {
      "clientId": "",
      "clientSecret": "",
      "domain": "",
      "redirectURL": ""
    },
    "seznam": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "strava": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "hubspot": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "line": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "atlassian": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "apple": {
      "teamId": "",
      "keyId": "",
      "privateKey": "",
      "redirectURL": "",
      "clientId": ""
    },
    "kick": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": ""
    },
    "livechat": {
      "clientId": "",
      "clientSecret": ""
    },
    "salesforce": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": "",
      "baseURL": "",
      "scope": ""
    },
    "slack": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": "",
      "scope": ""
    },
    "heroku": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": "",
      "scope": ""
    },
    "roblox": {
      "clientId": "",
      "clientSecret": "",
      "redirectURL": "",
      "scope": ""
    },
    "okta": {
      "clientId": "",
      "clientSecret": "",
      "domain": "",
      "audience": "",
      "scope": [],
      "redirectURL": ""
    },
    "ory": {
      "clientId": "",
      "clientSecret": "",
      "sdkURL": "",
      "redirectURL": "",
      "scope": [],
      "authorizationURL": "",
      "tokenURL": "",
      "userURL": ""
    },
    "shopifyCustomer": {
      "shopDomain": "",
      "clientId": "",
      "redirectURL": "",
      "scope": []
    }
  }
};
const envOptions = {
  prefix: "NITRO_",
  altPrefix: _inlineRuntimeConfig.nitro.envPrefix ?? process.env.NITRO_ENV_PREFIX ?? "_",
  envExpansion: _inlineRuntimeConfig.nitro.envExpansion ?? process.env.NITRO_ENV_EXPANSION ?? false
};
const _sharedRuntimeConfig = _deepFreeze(
  applyEnv(klona(_inlineRuntimeConfig), envOptions)
);
function useRuntimeConfig(event) {
  if (!event) {
    return _sharedRuntimeConfig;
  }
  if (event.context.nitro.runtimeConfig) {
    return event.context.nitro.runtimeConfig;
  }
  const runtimeConfig = klona(_inlineRuntimeConfig);
  applyEnv(runtimeConfig, envOptions);
  event.context.nitro.runtimeConfig = runtimeConfig;
  return runtimeConfig;
}
const _sharedAppConfig = _deepFreeze(klona(appConfig));
function useAppConfig(event) {
  {
    return _sharedAppConfig;
  }
}
function _deepFreeze(object) {
  const propNames = Object.getOwnPropertyNames(object);
  for (const name of propNames) {
    const value = object[name];
    if (value && typeof value === "object") {
      _deepFreeze(value);
    }
  }
  return Object.freeze(object);
}
new Proxy(/* @__PURE__ */ Object.create(null), {
  get: (_, prop) => {
    console.warn(
      "Please use `useRuntimeConfig()` instead of accessing config directly."
    );
    const runtimeConfig = useRuntimeConfig();
    if (prop in runtimeConfig) {
      return runtimeConfig[prop];
    }
    return void 0;
  }
});

const nitroAsyncContext = getContext("nitro-app", {
  asyncContext: true,
  AsyncLocalStorage: AsyncLocalStorage 
});

const config = useRuntimeConfig();
const _routeRulesMatcher = toRouteMatcher(
  createRouter({ routes: config.nitro.routeRules })
);
function createRouteRulesHandler(ctx) {
  return eventHandler((event) => {
    const routeRules = getRouteRules(event);
    if (routeRules.headers) {
      setHeaders(event, routeRules.headers);
    }
    if (routeRules.redirect) {
      let target = routeRules.redirect.to;
      if (target.endsWith("/**")) {
        let targetPath = event.path;
        const strpBase = routeRules.redirect._redirectStripBase;
        if (strpBase) {
          targetPath = withoutBase(targetPath, strpBase);
        }
        target = joinURL(target.slice(0, -3), targetPath);
      } else if (event.path.includes("?")) {
        const query = getQuery(event.path);
        target = withQuery(target, query);
      }
      return sendRedirect(event, target, routeRules.redirect.statusCode);
    }
    if (routeRules.proxy) {
      let target = routeRules.proxy.to;
      if (target.endsWith("/**")) {
        let targetPath = event.path;
        const strpBase = routeRules.proxy._proxyStripBase;
        if (strpBase) {
          targetPath = withoutBase(targetPath, strpBase);
        }
        target = joinURL(target.slice(0, -3), targetPath);
      } else if (event.path.includes("?")) {
        const query = getQuery(event.path);
        target = withQuery(target, query);
      }
      return proxyRequest(event, target, {
        fetch: ctx.localFetch,
        ...routeRules.proxy
      });
    }
  });
}
function getRouteRules(event) {
  event.context._nitro = event.context._nitro || {};
  if (!event.context._nitro.routeRules) {
    event.context._nitro.routeRules = getRouteRulesForPath(
      withoutBase(event.path.split("?")[0], useRuntimeConfig().app.baseURL)
    );
  }
  return event.context._nitro.routeRules;
}
function getRouteRulesForPath(path) {
  return defu({}, ..._routeRulesMatcher.matchAll(path).reverse());
}

function _captureError(error, type) {
  console.error(`[${type}]`, error);
  useNitroApp().captureError(error, { tags: [type] });
}
function trapUnhandledNodeErrors() {
  process.on(
    "unhandledRejection",
    (error) => _captureError(error, "unhandledRejection")
  );
  process.on(
    "uncaughtException",
    (error) => _captureError(error, "uncaughtException")
  );
}
function joinHeaders(value) {
  return Array.isArray(value) ? value.join(", ") : String(value);
}
function normalizeFetchResponse(response) {
  if (!response.headers.has("set-cookie")) {
    return response;
  }
  return new Response(response.body, {
    status: response.status,
    statusText: response.statusText,
    headers: normalizeCookieHeaders(response.headers)
  });
}
function normalizeCookieHeader(header = "") {
  return splitCookiesString(joinHeaders(header));
}
function normalizeCookieHeaders(headers) {
  const outgoingHeaders = new Headers();
  for (const [name, header] of headers) {
    if (name === "set-cookie") {
      for (const cookie of normalizeCookieHeader(header)) {
        outgoingHeaders.append("set-cookie", cookie);
      }
    } else {
      outgoingHeaders.set(name, joinHeaders(header));
    }
  }
  return outgoingHeaders;
}

function isJsonRequest(event) {
  if (hasReqHeader(event, "accept", "text/html")) {
    return false;
  }
  return hasReqHeader(event, "accept", "application/json") || hasReqHeader(event, "user-agent", "curl/") || hasReqHeader(event, "user-agent", "httpie/") || hasReqHeader(event, "sec-fetch-mode", "cors") || event.path.startsWith("/api/") || event.path.endsWith(".json");
}
function hasReqHeader(event, name, includes) {
  const value = getRequestHeader(event, name);
  return value && typeof value === "string" && value.toLowerCase().includes(includes);
}

const iframeStorageBridge = (nonce) => (
  /* js */
  `
(function() {
  const memoryStore = {};

  const NONCE = ${JSON.stringify(nonce)}
  
  const mockStorage = {
    getItem: function(key) {
      return memoryStore[key] !== undefined ? memoryStore[key] : null;
    },
    setItem: function(key, value) {
      memoryStore[key] = String(value);
      window.parent.postMessage({
        type: 'storage-set',
        key: key,
        value: String(value),
        nonce: NONCE
      }, '*');
    },
    removeItem: function(key) {
      delete memoryStore[key];
      window.parent.postMessage({
        type: 'storage-remove',
        key: key,
        nonce: NONCE
      }, '*');
    },
    clear: function() {
      for (const key in memoryStore) {
        delete memoryStore[key];
      }
      window.parent.postMessage({
        type: 'storage-clear',
        nonce: NONCE
      }, '*');
    },
    key: function(index) {
      const keys = Object.keys(memoryStore);
      return keys[index] !== undefined ? keys[index] : null;
    },
    get length() {
      return Object.keys(memoryStore).length;
    }
  };
  
  try {
    Object.defineProperty(window, 'localStorage', {
      value: mockStorage,
      writable: false,
      configurable: true
    });
  } catch (e) {
    window.localStorage = mockStorage;
  }
  
  window.addEventListener('message', function(event) {
    if (event.data.type === 'storage-sync-data' && event.data.nonce === NONCE) {
      const data = event.data.data;
      for (const key in data) {
        if (Object.prototype.hasOwnProperty.call(data, key)) {
          memoryStore[key] = data[key];
        }
      }
      if (typeof window.initTheme === 'function') {
        window.initTheme();
      }
      window.dispatchEvent(new Event('storage-ready'));
    }
  });
  
  window.parent.postMessage({ 
    type: 'storage-sync-request',
    nonce: NONCE
  }, '*');
})();
`
);
const parentStorageBridge = (nonce) => (
  /* js */
  `
(function() {
  const host = document.querySelector('nuxt-error-overlay');
  if (!host) return;
  
  // Wait for shadow root to be attached
  const checkShadow = setInterval(function() {
    if (host.shadowRoot) {
      clearInterval(checkShadow);
      const iframe = host.shadowRoot.getElementById('frame');
      if (!iframe) return;

      const NONCE = ${JSON.stringify(nonce)}
      
      window.addEventListener('message', function(event) {
        if (!event.data || event.data.nonce !== NONCE) return;
        
        const data = event.data;
        
        if (data.type === 'storage-set') {
          localStorage.setItem(data.key, data.value);
        } else if (data.type === 'storage-remove') {
          localStorage.removeItem(data.key);
        } else if (data.type === 'storage-clear') {
          localStorage.clear();
        } else if (data.type === 'storage-sync-request') {
          const allData = {};
          for (let i = 0; i < localStorage.length; i++) {
            const key = localStorage.key(i);
            allData[key] = localStorage.getItem(key);
          }
          iframe.contentWindow.postMessage({
            type: 'storage-sync-data',
            data: allData,
            nonce: NONCE
          }, '*');
        }
      });
    }
  }, 10);
})();
`
);
const errorCSS = (
  /* css */
  `
:host {
  --preview-width: 240px;
  --preview-height: 180px;
  --base-width: 1200px;
  --base-height: 900px;
  --z-base: 999999998;
  all: initial;
  display: contents;
}
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}
#frame {
  position: fixed;
  left: 0;
  top: 0;
  width: 100vw;
  height: 100vh;
  border: none;
  z-index: var(--z-base);
}
#frame[inert] {
  right: 5px;
  bottom: 5px;
  left: auto;
  top: auto;
  width: var(--base-width);
  height: var(--base-height);
  transform: scale(calc(240 / 1200));
  transform-origin: bottom right;
  overflow: hidden;
  border-radius: calc(1200 * 8px / 240);
}
#preview {
  position: fixed;
  right: 5px;
  bottom: 5px;
  width: var(--preview-width);
  height: var(--preview-height);
  overflow: hidden;
  border-radius: 8px;
  pointer-events: none;
  z-index: var(--z-base);
  background: white;
  display: none;
}
#frame:not([inert]) + #preview {
  display: block;
}
#toggle {
  position: fixed;
  right: 5px;
  bottom: 5px;
  width: var(--preview-width);
  height: var(--preview-height);
  background: none;
  border: 3px solid #00DC82;
  border-radius: 8px;
  cursor: pointer;
  opacity: 0.8;
  transition: opacity 0.2s, box-shadow 0.2s;
  z-index: calc(var(--z-base) + 1);
}
#toggle:hover,
#toggle:focus {
  opacity: 1;
  box-shadow: 0 0 20px rgba(0, 220, 130, 0.6);
}
#toggle:focus-visible {
  outline: 3px solid #00DC82;
  outline-offset: 3px;
  box-shadow: 0 0 24px rgba(0, 220, 130, 0.8);
}
@media (prefers-reduced-motion: reduce) {
  #toggle {
    transition: none;
  }
}
`
);
function webComponentScript(base64HTML, startMinimized) {
  return (
    /* js */
    `
  (function() {
    try {
      const host = document.querySelector('nuxt-error-overlay');
      if (!host) return;
      
      const shadow = host.attachShadow({ mode: 'open' });
      
      // Create elements
      const style = document.createElement('style');
      style.textContent = ${JSON.stringify(errorCSS)};
      
      const iframe = document.createElement('iframe');
      iframe.id = 'frame';
      iframe.src = 'data:text/html;base64,${base64HTML}';
      iframe.title = 'Detailed error stack trace';
      iframe.setAttribute('sandbox', 'allow-scripts allow-same-origin');
      
      const preview = document.createElement('div');
      preview.id = 'preview';
      
      const button = document.createElement('button');
      button.id = 'toggle';
      button.setAttribute('aria-expanded', 'true');
      button.setAttribute('type', 'button');
      button.innerHTML = '<span class="sr-only">Toggle detailed error view</span>';
      
      const liveRegion = document.createElement('div');
      liveRegion.setAttribute('role', 'status');
      liveRegion.setAttribute('aria-live', 'polite');
      liveRegion.className = 'sr-only';
      
      // Update preview snapshot
      function updatePreview() {
        try {
          let previewIframe = preview.querySelector('iframe');
          if (!previewIframe) {
            previewIframe = document.createElement('iframe');
            previewIframe.style.cssText = 'width: 1200px; height: 900px; transform: scale(0.2); transform-origin: top left; border: none;';
            previewIframe.setAttribute('sandbox', 'allow-scripts allow-same-origin');
            preview.appendChild(previewIframe);
          }
          
          const doctype = document.doctype ? '<!DOCTYPE ' + document.doctype.name + '>' : '';
          const cleanedHTML = document.documentElement.outerHTML
            .replace(/<nuxt-error-overlay[^>]*>.*?<\\/nuxt-error-overlay>/gs, '')
            .replace(/<script[^>]*>.*?<\\/script>/gs, '');
          
          const iframeDoc = previewIframe.contentDocument || previewIframe.contentWindow.document;
          iframeDoc.open();
          iframeDoc.write(doctype + cleanedHTML);
          iframeDoc.close();
        } catch (error) {
          console.error('Failed to update preview:', error);
        }
      }
      
      function toggleView() {
        const isMinimized = iframe.hasAttribute('inert');
        
        if (isMinimized) {
          updatePreview();
          iframe.removeAttribute('inert');
          button.setAttribute('aria-expanded', 'true');
          liveRegion.textContent = 'Showing detailed error view';
          setTimeout(function() {
            try { iframe.contentWindow.focus(); } catch {}
          }, 100);
        } else {
          iframe.setAttribute('inert', '');
          button.setAttribute('aria-expanded', 'false');
          liveRegion.textContent = 'Showing error page';
          button.focus();
        }
      }
      
      button.onclick = toggleView;
      
      document.addEventListener('keydown', function(e) {
        if ((e.key === 'Escape' || e.key === 'Esc') && !iframe.hasAttribute('inert')) {
          toggleView();
        }
      });
      
      // Append to shadow DOM
      shadow.appendChild(style);
      shadow.appendChild(liveRegion);
      shadow.appendChild(iframe);
      shadow.appendChild(preview);
      shadow.appendChild(button);
      
      if (${startMinimized}) {
        iframe.setAttribute('inert', '');
        button.setAttribute('aria-expanded', 'false');
      }
      
      // Initialize preview
      setTimeout(updatePreview, 100);
      
    } catch (error) {
      console.error('Failed to initialize Nuxt error overlay:', error);
    }
  })();
  `
  );
}
function generateErrorOverlayHTML(html, options) {
  const nonce = Array.from(crypto.getRandomValues(new Uint8Array(16)), (b) => b.toString(16).padStart(2, "0")).join("");
  const errorPage = html.replace("<head>", `<head><script>${iframeStorageBridge(nonce)}<\/script>`);
  const base64HTML = Buffer.from(errorPage, "utf8").toString("base64");
  return `
    <script>${parentStorageBridge(nonce)}<\/script>
    <nuxt-error-overlay></nuxt-error-overlay>
    <script>${webComponentScript(base64HTML, options?.startMinimized ?? false)}<\/script>
  `;
}

const errorHandler$0 = (async function errorhandler(error, event, { defaultHandler }) {
  if (event.handled || isJsonRequest(event)) {
    return;
  }
  const defaultRes = await defaultHandler(error, event, { json: true });
  const statusCode = error.statusCode || 500;
  if (statusCode === 404 && defaultRes.status === 302) {
    setResponseHeaders(event, defaultRes.headers);
    setResponseStatus(event, defaultRes.status, defaultRes.statusText);
    return send(event, JSON.stringify(defaultRes.body, null, 2));
  }
  if (typeof defaultRes.body !== "string" && Array.isArray(defaultRes.body.stack)) {
    defaultRes.body.stack = defaultRes.body.stack.join("\n");
  }
  const errorObject = defaultRes.body;
  const url = new URL(errorObject.url);
  errorObject.url = withoutBase(url.pathname, useRuntimeConfig(event).app.baseURL) + url.search + url.hash;
  errorObject.message ||= "Server Error";
  errorObject.data ||= error.data;
  errorObject.statusMessage ||= error.statusMessage;
  delete defaultRes.headers["content-type"];
  delete defaultRes.headers["content-security-policy"];
  setResponseHeaders(event, defaultRes.headers);
  const reqHeaders = getRequestHeaders(event);
  const isRenderingError = event.path.startsWith("/__nuxt_error") || !!reqHeaders["x-nuxt-error"];
  const res = isRenderingError ? null : await useNitroApp().localFetch(
    withQuery(joinURL(useRuntimeConfig(event).app.baseURL, "/__nuxt_error"), errorObject),
    {
      headers: { ...reqHeaders, "x-nuxt-error": "true" },
      redirect: "manual"
    }
  ).catch(() => null);
  if (event.handled) {
    return;
  }
  if (!res) {
    const { template } = await Promise.resolve().then(function () { return error500; });
    {
      errorObject.description = errorObject.message;
    }
    setResponseHeader(event, "Content-Type", "text/html;charset=UTF-8");
    return send(event, template(errorObject));
  }
  const html = await res.text();
  for (const [header, value] of res.headers.entries()) {
    if (header === "set-cookie") {
      appendResponseHeader(event, header, value);
      continue;
    }
    setResponseHeader(event, header, value);
  }
  setResponseStatus(event, res.status && res.status !== 200 ? res.status : defaultRes.status, res.statusText || defaultRes.statusText);
  if (!globalThis._importMeta_.test && typeof html === "string") {
    const prettyResponse = await defaultHandler(error, event, { json: false });
    return send(event, html.replace("</body>", `${generateErrorOverlayHTML(prettyResponse.body, { startMinimized: 300 <= statusCode && statusCode < 500 })}</body>`));
  }
  return send(event, html);
});

function defineNitroErrorHandler(handler) {
  return handler;
}

const errorHandler$1 = defineNitroErrorHandler(
  async function defaultNitroErrorHandler(error, event) {
    const res = await defaultHandler(error, event);
    if (!event.node?.res.headersSent) {
      setResponseHeaders(event, res.headers);
    }
    setResponseStatus(event, res.status, res.statusText);
    return send(
      event,
      typeof res.body === "string" ? res.body : JSON.stringify(res.body, null, 2)
    );
  }
);
async function defaultHandler(error, event, opts) {
  const isSensitive = error.unhandled || error.fatal;
  const statusCode = error.statusCode || 500;
  const statusMessage = error.statusMessage || "Server Error";
  const url = getRequestURL(event, { xForwardedHost: true, xForwardedProto: true });
  if (statusCode === 404) {
    const baseURL = "/";
    if (/^\/[^/]/.test(baseURL) && !url.pathname.startsWith(baseURL)) {
      const redirectTo = `${baseURL}${url.pathname.slice(1)}${url.search}`;
      return {
        status: 302,
        statusText: "Found",
        headers: { location: redirectTo },
        body: `Redirecting...`
      };
    }
  }
  await loadStackTrace(error).catch(consola.error);
  const youch = new Youch();
  if (isSensitive && !opts?.silent) {
    const tags = [error.unhandled && "[unhandled]", error.fatal && "[fatal]"].filter(Boolean).join(" ");
    const ansiError = await (await youch.toANSI(error)).replaceAll(process.cwd(), ".");
    consola.error(
      `[request error] ${tags} [${event.method}] ${url}

`,
      ansiError
    );
  }
  const useJSON = opts?.json || !getRequestHeader(event, "accept")?.includes("text/html");
  const headers = {
    "content-type": useJSON ? "application/json" : "text/html",
    // Prevent browser from guessing the MIME types of resources.
    "x-content-type-options": "nosniff",
    // Prevent error page from being embedded in an iframe
    "x-frame-options": "DENY",
    // Prevent browsers from sending the Referer header
    "referrer-policy": "no-referrer",
    // Disable the execution of any js
    "content-security-policy": "script-src 'self' 'unsafe-inline'; object-src 'none'; base-uri 'self';"
  };
  if (statusCode === 404 || !getResponseHeader(event, "cache-control")) {
    headers["cache-control"] = "no-cache";
  }
  const body = useJSON ? {
    error: true,
    url,
    statusCode,
    statusMessage,
    message: error.message,
    data: error.data,
    stack: error.stack?.split("\n").map((line) => line.trim())
  } : await youch.toHTML(error, {
    request: {
      url: url.href,
      method: event.method,
      headers: getRequestHeaders(event)
    }
  });
  return {
    status: statusCode,
    statusText: statusMessage,
    headers,
    body
  };
}
async function loadStackTrace(error) {
  if (!(error instanceof Error)) {
    return;
  }
  const parsed = await new ErrorParser().defineSourceLoader(sourceLoader).parse(error);
  const stack = error.message + "\n" + parsed.frames.map((frame) => fmtFrame(frame)).join("\n");
  Object.defineProperty(error, "stack", { value: stack });
  if (error.cause) {
    await loadStackTrace(error.cause).catch(consola.error);
  }
}
async function sourceLoader(frame) {
  if (!frame.fileName || frame.fileType !== "fs" || frame.type === "native") {
    return;
  }
  if (frame.type === "app") {
    const rawSourceMap = await readFile(`${frame.fileName}.map`, "utf8").catch(() => {
    });
    if (rawSourceMap) {
      const consumer = await new SourceMapConsumer(rawSourceMap);
      const originalPosition = consumer.originalPositionFor({ line: frame.lineNumber, column: frame.columnNumber });
      if (originalPosition.source && originalPosition.line) {
        frame.fileName = resolve(dirname(frame.fileName), originalPosition.source);
        frame.lineNumber = originalPosition.line;
        frame.columnNumber = originalPosition.column || 0;
      }
    }
  }
  const contents = await readFile(frame.fileName, "utf8").catch(() => {
  });
  return contents ? { contents } : void 0;
}
function fmtFrame(frame) {
  if (frame.type === "native") {
    return frame.raw;
  }
  const src = `${frame.fileName || ""}:${frame.lineNumber}:${frame.columnNumber})`;
  return frame.functionName ? `at ${frame.functionName} (${src}` : `at ${src}`;
}

const errorHandlers = [errorHandler$0, errorHandler$1];

async function errorHandler(error, event) {
  for (const handler of errorHandlers) {
    try {
      await handler(error, event, { defaultHandler });
      if (event.handled) {
        return; // Response handled
      }
    } catch(error) {
      // Handler itself thrown, log and continue
      console.error(error);
    }
  }
  // H3 will handle fallback
}

const script$1 = `
if (!window.__NUXT_DEVTOOLS_TIME_METRIC__) {
  Object.defineProperty(window, '__NUXT_DEVTOOLS_TIME_METRIC__', {
    value: {},
    enumerable: false,
    configurable: true,
  })
}
window.__NUXT_DEVTOOLS_TIME_METRIC__.appInit = Date.now()
`;

const _M9LpJilVcXk4b68dL_9V7Z_ArqaAQT4CmgEVKkd44 = (function(nitro) {
  nitro.hooks.hook("render:html", (htmlContext) => {
    htmlContext.head.push(`<script>${script$1}<\/script>`);
  });
});

const _IEVUMSdw9VzX0U84AT2_ckN8utdQLwtiTxlb4ORvNk = defineNitroPlugin((nitroApp) => {
  if (process.env.NUXT_OAUTH_FACEBOOK_CLIENT_ID && process.env.NUXT_OAUTH_FACEBOOK_CLIENT_SECRET || process.env.NUXT_OAUTH_INSTAGRAM_CLIENT_ID && process.env.NUXT_OAUTH_INSTAGRAM_CLIENT_SECRET) {
    nitroApp.hooks.hook("render:html", (html) => {
      html.head.unshift(`
      <script>
        if (window.location.hash === "#_=_"){
          history.replaceState
              ? history.replaceState(null, null, window.location.href.split("#")[0])
              : window.location.hash = "";
        }
      <\/script>
    `);
    });
  }
});

const rootDir = "/Users/mojito/Code/atidraw";

const appHead = {"meta":[{"name":"viewport","content":"width=device-width, initial-scale=1"},{"charset":"utf-8"}],"link":[],"style":[],"script":[],"noscript":[]};

const appRootTag = "div";

const appRootAttrs = {"id":"__nuxt","class":"isolate"};

const appTeleportTag = "div";

const appTeleportAttrs = {"id":"teleports"};

const appSpaLoaderTag = "div";

const appSpaLoaderAttrs = {"id":"__nuxt-loader"};

const appId = "nuxt-app";

const devReducers = {
  VNode: (data) => isVNode(data) ? { type: data.type, props: data.props } : void 0,
  URL: (data) => data instanceof URL ? data.toString() : void 0
};
const asyncContext = getContext("nuxt-dev", { asyncContext: true, AsyncLocalStorage });
const _ey9lThnP1vZr5FjKo9yU9M7oOsdOeHJ7C072wCMjFs = (nitroApp) => {
  const handler = nitroApp.h3App.handler;
  nitroApp.h3App.handler = (event) => {
    return asyncContext.callAsync({ logs: [], event }, () => handler(event));
  };
  onConsoleLog((_log) => {
    const ctx = asyncContext.tryUse();
    if (!ctx) {
      return;
    }
    const rawStack = captureRawStackTrace();
    if (!rawStack || rawStack.includes("runtime/vite-node.mjs")) {
      return;
    }
    const trace = [];
    let filename = "";
    for (const entry of parseRawStackTrace(rawStack)) {
      if (entry.source === globalThis._importMeta_.url) {
        continue;
      }
      if (EXCLUDE_TRACE_RE.test(entry.source)) {
        continue;
      }
      filename ||= entry.source.replace(withTrailingSlash(rootDir), "");
      trace.push({
        ...entry,
        source: entry.source.startsWith("file://") ? entry.source.replace("file://", "") : entry.source
      });
    }
    const log = {
      ..._log,
      // Pass along filename to allow the client to display more info about where log comes from
      filename,
      // Clean up file names in stack trace
      stack: trace
    };
    ctx.logs.push(log);
  });
  nitroApp.hooks.hook("afterResponse", () => {
    const ctx = asyncContext.tryUse();
    if (!ctx) {
      return;
    }
    return nitroApp.hooks.callHook("dev:ssr-logs", { logs: ctx.logs, path: ctx.event.path });
  });
  nitroApp.hooks.hook("render:html", (htmlContext) => {
    const ctx = asyncContext.tryUse();
    if (!ctx) {
      return;
    }
    try {
      const reducers = Object.assign(/* @__PURE__ */ Object.create(null), devReducers, ctx.event.context._payloadReducers);
      htmlContext.bodyAppend.unshift(`<script type="application/json" data-nuxt-logs="${appId}">${stringify(ctx.logs, reducers)}<\/script>`);
    } catch (e) {
      const shortError = e instanceof Error && "toString" in e ? ` Received \`${e.toString()}\`.` : "";
      console.warn(`[nuxt] Failed to stringify dev server logs.${shortError} You can define your own reducer/reviver for rich types following the instructions in https://nuxt.com/docs/4.x/api/composables/use-nuxt-app#payload.`);
    }
  });
};
const EXCLUDE_TRACE_RE = /\/node_modules\/(?:.*\/)?(?:nuxt|nuxt-nightly|nuxt-edge|nuxt3|consola|@vue)\/|core\/runtime\/nitro/;
function onConsoleLog(callback) {
  consola$1.addReporter({
    log(logObj) {
      callback(logObj);
    }
  });
  consola$1.wrapConsole();
}

const script = "\"use strict\";(()=>{const t=window,e=document.documentElement,c=[\"dark\",\"light\"],n=getStorageValue(\"localStorage\",\"nuxt-color-mode\")||\"system\";let i=n===\"system\"?u():n;const r=e.getAttribute(\"data-color-mode-forced\");r&&(i=r),l(i),t[\"__NUXT_COLOR_MODE__\"]={preference:n,value:i,getColorScheme:u,addColorScheme:l,removeColorScheme:d};function l(o){const s=\"\"+o+\"\",a=\"\";e.classList?e.classList.add(s):e.className+=\" \"+s,a&&e.setAttribute(\"data-\"+a,o)}function d(o){const s=\"\"+o+\"\",a=\"\";e.classList?e.classList.remove(s):e.className=e.className.replace(new RegExp(s,\"g\"),\"\"),a&&e.removeAttribute(\"data-\"+a)}function f(o){return t.matchMedia(\"(prefers-color-scheme\"+o+\")\")}function u(){if(t.matchMedia&&f(\"\").media!==\"not all\"){for(const o of c)if(f(\":\"+o).matches)return o}return\"light\"}})();function getStorageValue(t,e){switch(t){case\"localStorage\":return window.localStorage.getItem(e);case\"sessionStorage\":return window.sessionStorage.getItem(e);case\"cookie\":return getCookie(e);default:return null}}function getCookie(t){const c=(\"; \"+window.document.cookie).split(\"; \"+t+\"=\");if(c.length===2)return c.pop()?.split(\";\").shift()}";

const _RbtKAPooiysDsS8NUC0wAbU9iCWIK6cnaYS4JWN9Sp8 = (function(nitro) {
  nitro.hooks.hook("render:html", (htmlContext) => {
    htmlContext.head.push(`<script>${script}<\/script>`);
  });
});

const plugins = [
  _M9LpJilVcXk4b68dL_9V7Z_ArqaAQT4CmgEVKkd44,
_IEVUMSdw9VzX0U84AT2_ckN8utdQLwtiTxlb4ORvNk,
_ey9lThnP1vZr5FjKo9yU9M7oOsdOeHJ7C072wCMjFs,
_RbtKAPooiysDsS8NUC0wAbU9iCWIK6cnaYS4JWN9Sp8
];

const assets = {};

function readAsset (id) {
  const serverDir = dirname$1(fileURLToPath(globalThis._importMeta_.url));
  return promises.readFile(resolve$1(serverDir, assets[id].path))
}

const publicAssetBases = {"/_nuxt/builds/meta/":{"maxAge":31536000},"/_nuxt/builds/":{"maxAge":1}};

function isPublicAssetURL(id = '') {
  if (assets[id]) {
    return true
  }
  for (const base in publicAssetBases) {
    if (id.startsWith(base)) { return true }
  }
  return false
}

function getAsset (id) {
  return assets[id]
}

const METHODS = /* @__PURE__ */ new Set(["HEAD", "GET"]);
const EncodingMap = { gzip: ".gz", br: ".br" };
const _yyFlzC = eventHandler((event) => {
  if (event.method && !METHODS.has(event.method)) {
    return;
  }
  let id = decodePath(
    withLeadingSlash(withoutTrailingSlash(parseURL(event.path).pathname))
  );
  let asset;
  const encodingHeader = String(
    getRequestHeader(event, "accept-encoding") || ""
  );
  const encodings = [
    ...encodingHeader.split(",").map((e) => EncodingMap[e.trim()]).filter(Boolean).sort(),
    ""
  ];
  if (encodings.length > 1) {
    appendResponseHeader(event, "Vary", "Accept-Encoding");
  }
  for (const encoding of encodings) {
    for (const _id of [id + encoding, joinURL(id, "index.html" + encoding)]) {
      const _asset = getAsset(_id);
      if (_asset) {
        asset = _asset;
        id = _id;
        break;
      }
    }
  }
  if (!asset) {
    if (isPublicAssetURL(id)) {
      removeResponseHeader(event, "Cache-Control");
      throw createError({ statusCode: 404 });
    }
    return;
  }
  const ifNotMatch = getRequestHeader(event, "if-none-match") === asset.etag;
  if (ifNotMatch) {
    setResponseStatus(event, 304, "Not Modified");
    return "";
  }
  const ifModifiedSinceH = getRequestHeader(event, "if-modified-since");
  const mtimeDate = new Date(asset.mtime);
  if (ifModifiedSinceH && asset.mtime && new Date(ifModifiedSinceH) >= mtimeDate) {
    setResponseStatus(event, 304, "Not Modified");
    return "";
  }
  if (asset.type && !getResponseHeader(event, "Content-Type")) {
    setResponseHeader(event, "Content-Type", asset.type);
  }
  if (asset.etag && !getResponseHeader(event, "ETag")) {
    setResponseHeader(event, "ETag", asset.etag);
  }
  if (asset.mtime && !getResponseHeader(event, "Last-Modified")) {
    setResponseHeader(event, "Last-Modified", mtimeDate.toUTCString());
  }
  if (asset.encoding && !getResponseHeader(event, "Content-Encoding")) {
    setResponseHeader(event, "Content-Encoding", asset.encoding);
  }
  if (asset.size > 0 && !getResponseHeader(event, "Content-Length")) {
    setResponseHeader(event, "Content-Length", asset.size);
  }
  return readAsset(id);
});

const VueResolver = (_, value) => {
  return isRef(value) ? toValue(value) : value;
};

const headSymbol = "usehead";
// @__NO_SIDE_EFFECTS__
function vueInstall(head) {
  const plugin = {
    install(app) {
      app.config.globalProperties.$unhead = head;
      app.config.globalProperties.$head = head;
      app.provide(headSymbol, head);
    }
  };
  return plugin.install;
}

// @__NO_SIDE_EFFECTS__
function resolveUnrefHeadInput(input) {
  return walkResolver(input, VueResolver);
}

// @__NO_SIDE_EFFECTS__
function createHead(options = {}) {
  const head = createHead$1({
    ...options,
    propResolvers: [VueResolver]
  });
  head.install = vueInstall(head);
  return head;
}

const unheadOptions = {
  disableDefaults: true,
};

function createSSRContext(event) {
  const ssrContext = {
    url: event.path,
    event,
    runtimeConfig: useRuntimeConfig(event),
    noSSR: event.context.nuxt?.noSSR || (false),
    head: createHead(unheadOptions),
    error: false,
    nuxt: void 0,
    /* NuxtApp */
    payload: {},
    _payloadReducers: /* @__PURE__ */ Object.create(null),
    modules: /* @__PURE__ */ new Set()
  };
  return ssrContext;
}
function setSSRError(ssrContext, error) {
  ssrContext.error = true;
  ssrContext.payload = { error };
  ssrContext.url = error.url;
}

function buildAssetsDir() {
  return useRuntimeConfig().app.buildAssetsDir;
}
function buildAssetsURL(...path) {
  return joinRelativeURL(publicAssetsURL(), buildAssetsDir(), ...path);
}
function publicAssetsURL(...path) {
  const app = useRuntimeConfig().app;
  const publicBase = app.cdnURL || app.baseURL;
  return path.length ? joinRelativeURL(publicBase, ...path) : publicBase;
}

const APP_ROOT_OPEN_TAG = `<${appRootTag}${propsToString(appRootAttrs)}>`;
const APP_ROOT_CLOSE_TAG = `</${appRootTag}>`;
const getServerEntry = () => import('file:///Users/mojito/Code/atidraw/.nuxt//dist/server/server.mjs').then((r) => r.default || r);
const getClientManifest = () => import('file:///Users/mojito/Code/atidraw/.nuxt//dist/server/client.manifest.mjs').then((r) => r.default || r).then((r) => typeof r === "function" ? r() : r);
const getSSRRenderer = lazyCachedFunction(async () => {
  const createSSRApp = await getServerEntry();
  if (!createSSRApp) {
    throw new Error("Server bundle is not available");
  }
  const precomputed = void 0 ;
  const renderer = createRenderer(createSSRApp, {
    precomputed,
    manifest: await getClientManifest() ,
    renderToString: renderToString$1,
    buildAssetsURL
  });
  async function renderToString$1(input, context) {
    const html = await renderToString(input, context);
    if (process$1.env.NUXT_VITE_NODE_OPTIONS) {
      renderer.rendererContext.updateManifest(await getClientManifest());
    }
    return APP_ROOT_OPEN_TAG + html + APP_ROOT_CLOSE_TAG;
  }
  return renderer;
});
const getSPARenderer = lazyCachedFunction(async () => {
  const precomputed = void 0 ;
  const spaTemplate = await Promise.resolve().then(function () { return _virtual__spaTemplate; }).then((r) => r.template).catch(() => "").then((r) => {
    {
      const APP_SPA_LOADER_OPEN_TAG = `<${appSpaLoaderTag}${propsToString(appSpaLoaderAttrs)}>`;
      const APP_SPA_LOADER_CLOSE_TAG = `</${appSpaLoaderTag}>`;
      const appTemplate = APP_ROOT_OPEN_TAG + APP_ROOT_CLOSE_TAG;
      const loaderTemplate = r ? APP_SPA_LOADER_OPEN_TAG + r + APP_SPA_LOADER_CLOSE_TAG : "";
      return appTemplate + loaderTemplate;
    }
  });
  const renderer = createRenderer(() => () => {
  }, {
    precomputed,
    manifest: await getClientManifest() ,
    renderToString: () => spaTemplate,
    buildAssetsURL
  });
  const result = await renderer.renderToString({});
  const renderToString = (ssrContext) => {
    const config = useRuntimeConfig(ssrContext.event);
    ssrContext.modules ||= /* @__PURE__ */ new Set();
    ssrContext.payload.serverRendered = false;
    ssrContext.config = {
      public: config.public,
      app: config.app
    };
    return Promise.resolve(result);
  };
  return {
    rendererContext: renderer.rendererContext,
    renderToString
  };
});
function lazyCachedFunction(fn) {
  let res = null;
  return () => {
    if (res === null) {
      res = fn().catch((err) => {
        res = null;
        throw err;
      });
    }
    return res;
  };
}
function getRenderer(ssrContext) {
  return ssrContext.noSSR ? getSPARenderer() : getSSRRenderer();
}
const getSSRStyles = lazyCachedFunction(() => Promise.resolve().then(function () { return styles$1; }).then((r) => r.default || r));

async function renderInlineStyles(usedModules) {
  const styleMap = await getSSRStyles();
  const inlinedStyles = /* @__PURE__ */ new Set();
  for (const mod of usedModules) {
    if (mod in styleMap && styleMap[mod]) {
      for (const style of await styleMap[mod]()) {
        inlinedStyles.add(style);
      }
    }
  }
  return Array.from(inlinedStyles).map((style) => ({ innerHTML: style }));
}

const ROOT_NODE_REGEX = new RegExp(`^<${appRootTag}[^>]*>([\\s\\S]*)<\\/${appRootTag}>$`);
function getServerComponentHTML(body) {
  const match = body.match(ROOT_NODE_REGEX);
  return match?.[1] || body;
}
const SSR_SLOT_TELEPORT_MARKER = /^uid=([^;]*);slot=(.*)$/;
const SSR_CLIENT_TELEPORT_MARKER = /^uid=([^;]*);client=(.*)$/;
const SSR_CLIENT_SLOT_MARKER = /^island-slot=([^;]*);(.*)$/;
function getSlotIslandResponse(ssrContext) {
  if (!ssrContext.islandContext || !Object.keys(ssrContext.islandContext.slots).length) {
    return void 0;
  }
  const response = {};
  for (const [name, slot] of Object.entries(ssrContext.islandContext.slots)) {
    response[name] = {
      ...slot,
      fallback: ssrContext.teleports?.[`island-fallback=${name}`]
    };
  }
  return response;
}
function getClientIslandResponse(ssrContext) {
  if (!ssrContext.islandContext || !Object.keys(ssrContext.islandContext.components).length) {
    return void 0;
  }
  const response = {};
  for (const [clientUid, component] of Object.entries(ssrContext.islandContext.components)) {
    const html = ssrContext.teleports?.[clientUid]?.replaceAll("<!--teleport start anchor-->", "") || "";
    response[clientUid] = {
      ...component,
      html,
      slots: getComponentSlotTeleport(clientUid, ssrContext.teleports ?? {})
    };
  }
  return response;
}
function getComponentSlotTeleport(clientUid, teleports) {
  const entries = Object.entries(teleports);
  const slots = {};
  for (const [key, value] of entries) {
    const match = key.match(SSR_CLIENT_SLOT_MARKER);
    if (match) {
      const [, id, slot] = match;
      if (!slot || clientUid !== id) {
        continue;
      }
      slots[slot] = value;
    }
  }
  return slots;
}
function replaceIslandTeleports(ssrContext, html) {
  const { teleports, islandContext } = ssrContext;
  if (islandContext || !teleports) {
    return html;
  }
  for (const key in teleports) {
    const matchClientComp = key.match(SSR_CLIENT_TELEPORT_MARKER);
    if (matchClientComp) {
      const [, uid, clientId] = matchClientComp;
      if (!uid || !clientId) {
        continue;
      }
      html = html.replace(new RegExp(` data-island-uid="${uid}" data-island-component="${clientId}"[^>]*>`), (full) => {
        return full + teleports[key];
      });
      continue;
    }
    const matchSlot = key.match(SSR_SLOT_TELEPORT_MARKER);
    if (matchSlot) {
      const [, uid, slot] = matchSlot;
      if (!uid || !slot) {
        continue;
      }
      html = html.replace(new RegExp(` data-island-uid="${uid}" data-island-slot="${slot}"[^>]*>`), (full) => {
        return full + teleports[key];
      });
    }
  }
  return html;
}

const ISLAND_SUFFIX_RE = /\.json(?:\?.*)?$/;
const _SxA8c9 = defineEventHandler(async (event) => {
  const nitroApp = useNitroApp();
  setResponseHeaders(event, {
    "content-type": "application/json;charset=utf-8",
    "x-powered-by": "Nuxt"
  });
  const islandContext = await getIslandContext(event);
  const ssrContext = {
    ...createSSRContext(event),
    islandContext,
    noSSR: false,
    url: islandContext.url
  };
  const renderer = await getSSRRenderer();
  const renderResult = await renderer.renderToString(ssrContext).catch(async (err) => {
    await ssrContext.nuxt?.hooks.callHook("app:error", err);
    throw err;
  });
  if (ssrContext.payload?.error) {
    throw ssrContext.payload.error;
  }
  const inlinedStyles = await renderInlineStyles(ssrContext.modules ?? []);
  await ssrContext.nuxt?.hooks.callHook("app:rendered", { ssrContext, renderResult });
  if (inlinedStyles.length) {
    ssrContext.head.push({ style: inlinedStyles });
  }
  {
    const { styles } = getRequestDependencies(ssrContext, renderer.rendererContext);
    const link = [];
    for (const resource of Object.values(styles)) {
      if ("inline" in getQuery(resource.file)) {
        continue;
      }
      if (resource.file.includes("scoped") && !resource.file.includes("pages/")) {
        link.push({ rel: "stylesheet", href: renderer.rendererContext.buildAssetsURL(resource.file), crossorigin: "" });
      }
    }
    if (link.length) {
      ssrContext.head.push({ link }, { mode: "server" });
    }
  }
  const islandHead = {};
  for (const entry of ssrContext.head.entries.values()) {
    for (const [key, value] of Object.entries(resolveUnrefHeadInput(entry.input))) {
      const currentValue = islandHead[key];
      if (Array.isArray(currentValue)) {
        currentValue.push(...value);
      }
      islandHead[key] = value;
    }
  }
  const islandResponse = {
    id: islandContext.id,
    head: islandHead,
    html: getServerComponentHTML(renderResult.html),
    components: getClientIslandResponse(ssrContext),
    slots: getSlotIslandResponse(ssrContext)
  };
  await nitroApp.hooks.callHook("render:island", islandResponse, { event, islandContext });
  return islandResponse;
});
async function getIslandContext(event) {
  let url = event.path || "";
  const componentParts = url.substring("/__nuxt_island".length + 1).replace(ISLAND_SUFFIX_RE, "").split("_");
  const hashId = componentParts.length > 1 ? componentParts.pop() : void 0;
  const componentName = componentParts.join("_");
  const context = event.method === "GET" ? getQuery$1(event) : await readBody(event);
  const ctx = {
    url: "/",
    ...context,
    id: hashId,
    name: componentName,
    props: destr$1(context.props) || {},
    slots: {},
    components: {}
  };
  return ctx;
}

function defineNitroPlugin(def) {
  return def;
}

function defineRenderHandler(render) {
  const runtimeConfig = useRuntimeConfig();
  return eventHandler(async (event) => {
    const nitroApp = useNitroApp();
    const ctx = { event, render, response: void 0 };
    await nitroApp.hooks.callHook("render:before", ctx);
    if (!ctx.response) {
      if (event.path === `${runtimeConfig.app.baseURL}favicon.ico`) {
        setResponseHeader(event, "Content-Type", "image/x-icon");
        return send(
          event,
          "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
        );
      }
      ctx.response = await ctx.render(event);
      if (!ctx.response) {
        const _currentStatus = getResponseStatus(event);
        setResponseStatus(event, _currentStatus === 200 ? 500 : _currentStatus);
        return send(
          event,
          "No response returned from render handler: " + event.path
        );
      }
    }
    await nitroApp.hooks.callHook("render:response", ctx.response, ctx);
    if (ctx.response.headers) {
      setResponseHeaders(event, ctx.response.headers);
    }
    if (ctx.response.statusCode || ctx.response.statusMessage) {
      setResponseStatus(
        event,
        ctx.response.statusCode,
        ctx.response.statusMessage
      );
    }
    return ctx.response.body;
  });
}

const scheduledTasks = false;

const tasks = {
  
};

const __runningTasks__ = {};
async function runTask(name, {
  payload = {},
  context = {}
} = {}) {
  if (__runningTasks__[name]) {
    return __runningTasks__[name];
  }
  if (!(name in tasks)) {
    throw createError({
      message: `Task \`${name}\` is not available!`,
      statusCode: 404
    });
  }
  if (!tasks[name].resolve) {
    throw createError({
      message: `Task \`${name}\` is not implemented!`,
      statusCode: 501
    });
  }
  const handler = await tasks[name].resolve();
  const taskEvent = { name, payload, context };
  __runningTasks__[name] = handler.run(taskEvent);
  try {
    const res = await __runningTasks__[name];
    return res;
  } finally {
    delete __runningTasks__[name];
  }
}

const OAUTH_COOKIE_MAX_AGE = 60 * 10;
function getOAuthRedirectURL(event) {
  const requestURL = getRequestURL(event);
  return `${requestURL.protocol}//${requestURL.host}${requestURL.pathname}`;
}
async function requestAccessToken(url, options) {
  const headers = {
    "Content-Type": "application/x-www-form-urlencoded",
    ...options.headers
  };
  const body = headers["Content-Type"] === "application/x-www-form-urlencoded" ? new URLSearchParams(
    options.body || options.params || {}
  ).toString() : options.body;
  return $fetch(url, {
    method: "POST",
    headers,
    body
  }).catch((error) => {
    if (error instanceof FetchError && error.status === 401) {
      return error.data;
    }
    throw error;
  });
}
function handleAccessTokenErrorResponse(event, oauthProvider, oauthError, onError) {
  const message = `${upperFirst(oauthProvider)} login failed: ${oauthError.error_description || oauthError.error || "Unknown error"}`;
  const error = createError({
    statusCode: 401,
    message,
    data: oauthError
  });
  if (!onError) throw error;
  return onError(event, error);
}
function handleMissingConfiguration(event, provider, missingKeys, onError) {
  const environmentVariables = missingKeys.map((key) => `NUXT_OAUTH_${provider.toUpperCase()}_${snakeCase(key).toUpperCase()}`);
  const error = createError({
    statusCode: 500,
    message: `Missing ${environmentVariables.join(" or ")} env ${missingKeys.length > 1 ? "variables" : "variable"}.`
  });
  if (!onError) throw error;
  return onError(event, error);
}
function handleInvalidState(event, provider, onError) {
  const message = `${upperFirst(provider)} login failed: state mismatch`;
  const error = createError({
    statusCode: 500,
    message
  });
  if (!onError) throw error;
  return onError(event, error);
}
function encodeBase64Url(input) {
  return btoa(String.fromCharCode.apply(null, input)).replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/g, "");
}
function getRandomBytes(size = 32) {
  return getRandomValues(new Uint8Array(size));
}
async function handleState(event) {
  const query = getQuery$1(event);
  if (query.state) {
    const state2 = getCookie(event, "nuxt-auth-state");
    deleteCookie(event, "nuxt-auth-state");
    return state2;
  }
  const state = encodeBase64Url(getRandomBytes(8));
  setCookie(event, "nuxt-auth-state", state, {
    httpOnly: true,
    secure: false,
    sameSite: "lax",
    maxAge: OAUTH_COOKIE_MAX_AGE,
    path: "/"
  });
  return state;
}

function defineOAuthGitHubEventHandler({ config, onSuccess, onError }) {
  return eventHandler(async (event) => {
    config = defu$1(config, useRuntimeConfig(event).oauth?.github, {
      authorizationURL: "https://github.com/login/oauth/authorize",
      tokenURL: "https://github.com/login/oauth/access_token",
      apiURL: "https://api.github.com",
      authorizationParams: {}
    });
    const query = getQuery$1(event);
    if (query.error) {
      const error = createError({
        statusCode: 401,
        message: `GitHub login failed: ${query.error || "Unknown error"}`,
        data: query
      });
      if (!onError) throw error;
      return onError(event, error);
    }
    if (!config.clientId || !config.clientSecret) {
      return handleMissingConfiguration(event, "github", ["clientId", "clientSecret"], onError);
    }
    const redirectURL = config.redirectURL || getOAuthRedirectURL(event);
    const state = await handleState(event);
    if (!query.code) {
      config.scope = config.scope || [];
      if (config.emailRequired && !config.scope.includes("user:email")) {
        config.scope.push("user:email");
      }
      return sendRedirect(
        event,
        withQuery(config.authorizationURL, {
          client_id: config.clientId,
          redirect_uri: redirectURL,
          scope: config.scope.join(" "),
          state,
          ...config.authorizationParams
        })
      );
    }
    if (query.state !== state) {
      return handleInvalidState(event, "github", onError);
    }
    const tokens = await requestAccessToken(config.tokenURL, {
      body: {
        grant_type: "authorization_code",
        client_id: config.clientId,
        client_secret: config.clientSecret,
        redirect_uri: redirectURL,
        code: query.code
      }
    });
    if (tokens.error) {
      return handleAccessTokenErrorResponse(event, "github", tokens, onError);
    }
    const accessToken = tokens.access_token;
    const user = await $fetch(`${config.apiURL}/user`, {
      headers: {
        "User-Agent": `Github-OAuth-${config.clientId}`,
        "Authorization": `token ${accessToken}`
      }
    });
    if (!user.email && config.emailRequired) {
      const emails = await $fetch(`${config.apiURL}/user/emails`, {
        headers: {
          "User-Agent": `Github-OAuth-${config.clientId}`,
          "Authorization": `token ${accessToken}`
        }
      });
      const primaryEmail = emails.find((email) => email.primary);
      if (!primaryEmail) {
        const error = createError({
          statusCode: 500,
          message: "Could not get GitHub user email",
          data: tokens
        });
        if (!onError) throw error;
        return onError(event, error);
      }
      user.email = primaryEmail.email;
      user.email_verified = primaryEmail.verified;
    }
    return onSuccess(event, {
      user,
      tokens
    });
  });
}

function defineOAuthGoogleEventHandler({
  config,
  onSuccess,
  onError
}) {
  return eventHandler(async (event) => {
    config = defu$1(config, useRuntimeConfig(event).oauth?.google, {
      authorizationURL: "https://accounts.google.com/o/oauth2/v2/auth",
      tokenURL: "https://oauth2.googleapis.com/token",
      userURL: "https://www.googleapis.com/oauth2/v3/userinfo",
      authorizationParams: {}
    });
    const query = getQuery$1(event);
    if (!config.clientId || !config.clientSecret) {
      return handleMissingConfiguration(event, "google", ["clientId", "clientSecret"], onError);
    }
    const redirectURL = config.redirectURL || getOAuthRedirectURL(event);
    if (!query.code) {
      config.scope = config.scope || ["email", "profile"];
      return sendRedirect(
        event,
        withQuery(config.authorizationURL, {
          response_type: "code",
          client_id: config.clientId,
          redirect_uri: redirectURL,
          scope: config.scope.join(" "),
          state: query.state || "",
          ...config.authorizationParams
        })
      );
    }
    const tokens = await requestAccessToken(config.tokenURL, {
      body: {
        grant_type: "authorization_code",
        code: query.code,
        client_id: config.clientId,
        client_secret: config.clientSecret,
        redirect_uri: redirectURL
      }
    });
    if (tokens.error) {
      return handleAccessTokenErrorResponse(event, "google", tokens, onError);
    }
    const accessToken = tokens.access_token;
    const user = await $fetch(
      config.userURL,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      }
    );
    return onSuccess(event, {
      tokens,
      user
    });
  });
}

const sessionHooks = createHooks();
async function getUserSession(event) {
  const session = await _useSession(event);
  return {
    ...session.data,
    id: session.id
  };
}
async function setUserSession(event, data, config) {
  const session = await _useSession(event, config);
  await session.update(defu$1(data, session.data));
  return session.data;
}
async function clearUserSession(event, config) {
  const session = await _useSession(event, config);
  await sessionHooks.callHookParallel("clear", session.data, event);
  await session.clear();
  return true;
}
async function requireUserSession(event, opts = {}) {
  const userSession = await getUserSession(event);
  if (!userSession.user) {
    if (isEvent(event)) {
      throw createError({
        statusCode: opts.statusCode || 401,
        message: opts.message || "Unauthorized"
      });
    } else {
      throw new Response(opts.message || "Unauthorized", {
        status: opts.statusCode || 401
      });
    }
  }
  return userSession;
}
let sessionConfig;
function _useSession(event, config = {}) {
  if (!sessionConfig) {
    const runtimeConfig = useRuntimeConfig(isEvent(event) ? event : void 0);
    const envSessionPassword = `${runtimeConfig.nitro?.envPrefix || "NUXT_"}SESSION_PASSWORD`;
    sessionConfig = defu$1({ password: process.env[envSessionPassword] }, runtimeConfig.session);
    if (!sessionConfig.password) {
      console.error(`[nuxt-auth-utils] ${envSessionPassword} environment variable or runtimeConfig.session.password was not set.`);
    }
  }
  const finalConfig = defu$1(config, sessionConfig);
  return useSession(event, finalConfig);
}

async function apiRequest(event, url, options = {}) {
  var _a, _b;
  const config = useRuntimeConfig();
  const apiBase = config.public.apiBase || "http://localhost:18888";
  const headers = {
    "Content-Type": "application/json",
    "Accept": "*/*",
    ...options.headers
  };
  let token = options.token;
  if (!token && event) {
    try {
      const session = await getUserSession(event);
      token = session == null ? void 0 : session.token;
    } catch {
    }
  }
  if (options.requireAuth && !token) {
    const error = new Error("\u9700\u8981\u767B\u5F55\u8BA4\u8BC1");
    error.statusCode = 401;
    throw error;
  }
  if (token) {
    headers.Authorization = `Bearer ${token}`;
  }
  try {
    const response = await $fetch(url, {
      method: options.method || "GET",
      baseURL: apiBase,
      headers,
      body: options.body
    });
    if (!response || typeof response !== "object" || !("code" in response)) {
      throw new Error("\u65E0\u6548\u7684API\u54CD\u5E94\u683C\u5F0F");
    }
    if (response.code !== 200) {
      throw new Error(response.msg || "\u8BF7\u6C42\u5931\u8D25");
    }
    return response;
  } catch (error) {
    const apiError = error;
    if (apiError.statusCode) {
      const err2 = new Error(((_a = apiError.data) == null ? void 0 : _a.message) || ((_b = apiError.data) == null ? void 0 : _b.msg) || apiError.message || "\u8BF7\u6C42\u5931\u8D25");
      err2.statusCode = apiError.statusCode;
      err2.data = apiError.data;
      throw err2;
    }
    if (apiError.data) {
      const err2 = new Error(apiError.data.message || apiError.data.msg || "\u8BF7\u6C42\u5931\u8D25");
      err2.statusCode = apiError.status || apiError.statusCode || 500;
      err2.data = apiError.data;
      throw err2;
    }
    const err = new Error(apiError.message || "\u7F51\u7EDC\u8BF7\u6C42\u5931\u8D25");
    err.statusCode = apiError.statusCode || 500;
    throw err;
  }
}

const warnOnceSet = /* @__PURE__ */ new Set();
const DEFAULT_ENDPOINT = "https://api.iconify.design";
const _4PhbgL = defineCachedEventHandler(async (event) => {
  const url = getRequestURL(event);
  if (!url)
    return createError({ status: 400, message: "Invalid icon request" });
  const options = useAppConfig().icon;
  const collectionName = event.context.params?.collection?.replace(/\.json$/, "");
  const collection = collectionName ? await collections[collectionName]?.() : null;
  const apiEndPoint = options.iconifyApiEndpoint || DEFAULT_ENDPOINT;
  const icons = url.searchParams.get("icons")?.split(",");
  if (collection) {
    if (icons?.length) {
      const data = getIcons(
        collection,
        icons
      );
      consola$1.debug(`[Icon] serving ${(icons || []).map((i) => "`" + collectionName + ":" + i + "`").join(",")} from bundled collection`);
      return data;
    }
  } else {
    if (collectionName && !warnOnceSet.has(collectionName) && apiEndPoint === DEFAULT_ENDPOINT) {
      consola$1.warn([
        `[Icon] Collection \`${collectionName}\` is not found locally`,
        `We suggest to install it via \`npm i -D @iconify-json/${collectionName}\` to provide the best end-user experience.`
      ].join("\n"));
      warnOnceSet.add(collectionName);
    }
  }
  if (options.fallbackToApi === true || options.fallbackToApi === "server-only") {
    const apiUrl = new URL("./" + basename(url.pathname) + url.search, apiEndPoint);
    consola$1.debug(`[Icon] fetching ${(icons || []).map((i) => "`" + collectionName + ":" + i + "`").join(",")} from iconify api`);
    if (apiUrl.host !== new URL(apiEndPoint).host) {
      return createError({ status: 400, message: "Invalid icon request" });
    }
    try {
      const data = await $fetch(apiUrl.href);
      return data;
    } catch (e) {
      consola$1.error(e);
      if (e.status === 404)
        return createError({ status: 404 });
      else
        return createError({ status: 500, message: "Failed to fetch fallback icon" });
    }
  }
  return createError({ status: 404 });
}, {
  group: "nuxt",
  name: "icon",
  getKey(event) {
    const collection = event.context.params?.collection?.replace(/\.json$/, "") || "unknown";
    const icons = String(getQuery$1(event).icons || "");
    return `${collection}_${icons.split(",")[0]}_${icons.length}_${hash$1(icons)}`;
  },
  swr: true,
  maxAge: 60 * 60 * 24 * 7
  // 1 week
});

const _BrR7TB = eventHandler(async (event) => {
  await clearUserSession(event);
  return { loggedOut: true };
});

const _iCqWxD = eventHandler(async (event) => {
  const session = await getUserSession(event);
  if (Object.keys(session).length > 0) {
    await sessionHooks.callHookParallel("fetch", session, event);
  }
  const { secure, ...data } = session;
  return data;
});

const _lazy_dsk5t0 = () => Promise.resolve().then(function () { return login_post$1; });
const _lazy_aoZZPm = () => Promise.resolve().then(function () { return register_post$1; });
const _lazy_joMom8 = () => Promise.resolve().then(function () { return drawings_get$1; });
const _lazy_KxhBLh = () => Promise.resolve().then(function () { return feed_get$1; });
const _lazy_2dPegD = () => Promise.resolve().then(function () { return generate_post$1; });
const _lazy_Gd5fFn = () => Promise.resolve().then(function () { return upload_post$1; });
const _lazy_JX8oIJ = () => Promise.resolve().then(function () { return anonymous_get$1; });
const _lazy_U1ZcoZ = () => Promise.resolve().then(function () { return github_get$1; });
const _lazy_JLw3dF = () => Promise.resolve().then(function () { return google_get$1; });
const _lazy_b6MkJD = () => Promise.resolve().then(function () { return ____pathname__get$1; });
const _lazy_nonZvk = () => Promise.resolve().then(function () { return renderer$1; });

const handlers = [
  { route: '', handler: _yyFlzC, lazy: false, middleware: true, method: undefined },
  { route: '/api/auth/login', handler: _lazy_dsk5t0, lazy: true, middleware: false, method: "post" },
  { route: '/api/auth/register', handler: _lazy_aoZZPm, lazy: true, middleware: false, method: "post" },
  { route: '/api/drawings', handler: _lazy_joMom8, lazy: true, middleware: false, method: "get" },
  { route: '/api/feed', handler: _lazy_KxhBLh, lazy: true, middleware: false, method: "get" },
  { route: '/api/generate', handler: _lazy_2dPegD, lazy: true, middleware: false, method: "post" },
  { route: '/api/upload', handler: _lazy_Gd5fFn, lazy: true, middleware: false, method: "post" },
  { route: '/auth/anonymous', handler: _lazy_JX8oIJ, lazy: true, middleware: false, method: "get" },
  { route: '/auth/github', handler: _lazy_U1ZcoZ, lazy: true, middleware: false, method: "get" },
  { route: '/auth/google', handler: _lazy_JLw3dF, lazy: true, middleware: false, method: "get" },
  { route: '/drawings/**:pathname', handler: _lazy_b6MkJD, lazy: true, middleware: false, method: "get" },
  { route: '/__nuxt_error', handler: _lazy_nonZvk, lazy: true, middleware: false, method: undefined },
  { route: '/__nuxt_island/**', handler: _SxA8c9, lazy: false, middleware: false, method: undefined },
  { route: '/api/_nuxt_icon/:collection', handler: _4PhbgL, lazy: false, middleware: false, method: undefined },
  { route: '/api/_auth/session', handler: _BrR7TB, lazy: false, middleware: false, method: "delete" },
  { route: '/api/_auth/session', handler: _iCqWxD, lazy: false, middleware: false, method: "get" },
  { route: '/**', handler: _lazy_nonZvk, lazy: true, middleware: false, method: undefined }
];

function createNitroApp() {
  const config = useRuntimeConfig();
  const hooks = createHooks();
  const captureError = (error, context = {}) => {
    const promise = hooks.callHookParallel("error", error, context).catch((error_) => {
      console.error("Error while capturing another error", error_);
    });
    if (context.event && isEvent(context.event)) {
      const errors = context.event.context.nitro?.errors;
      if (errors) {
        errors.push({ error, context });
      }
      if (context.event.waitUntil) {
        context.event.waitUntil(promise);
      }
    }
  };
  const h3App = createApp({
    debug: destr(true),
    onError: (error, event) => {
      captureError(error, { event, tags: ["request"] });
      return errorHandler(error, event);
    },
    onRequest: async (event) => {
      event.context.nitro = event.context.nitro || { errors: [] };
      const fetchContext = event.node.req?.__unenv__;
      if (fetchContext?._platform) {
        event.context = {
          _platform: fetchContext?._platform,
          // #3335
          ...fetchContext._platform,
          ...event.context
        };
      }
      if (!event.context.waitUntil && fetchContext?.waitUntil) {
        event.context.waitUntil = fetchContext.waitUntil;
      }
      event.fetch = (req, init) => fetchWithEvent(event, req, init, { fetch: localFetch });
      event.$fetch = (req, init) => fetchWithEvent(event, req, init, {
        fetch: $fetch
      });
      event.waitUntil = (promise) => {
        if (!event.context.nitro._waitUntilPromises) {
          event.context.nitro._waitUntilPromises = [];
        }
        event.context.nitro._waitUntilPromises.push(promise);
        if (event.context.waitUntil) {
          event.context.waitUntil(promise);
        }
      };
      event.captureError = (error, context) => {
        captureError(error, { event, ...context });
      };
      await nitroApp$1.hooks.callHook("request", event).catch((error) => {
        captureError(error, { event, tags: ["request"] });
      });
    },
    onBeforeResponse: async (event, response) => {
      await nitroApp$1.hooks.callHook("beforeResponse", event, response).catch((error) => {
        captureError(error, { event, tags: ["request", "response"] });
      });
    },
    onAfterResponse: async (event, response) => {
      await nitroApp$1.hooks.callHook("afterResponse", event, response).catch((error) => {
        captureError(error, { event, tags: ["request", "response"] });
      });
    }
  });
  const router = createRouter$1({
    preemptive: true
  });
  const nodeHandler = toNodeListener(h3App);
  const localCall = (aRequest) => callNodeRequestHandler(
    nodeHandler,
    aRequest
  );
  const localFetch = (input, init) => {
    if (!input.toString().startsWith("/")) {
      return globalThis.fetch(input, init);
    }
    return fetchNodeRequestHandler(
      nodeHandler,
      input,
      init
    ).then((response) => normalizeFetchResponse(response));
  };
  const $fetch = createFetch({
    fetch: localFetch,
    Headers: Headers$1,
    defaults: { baseURL: config.app.baseURL }
  });
  globalThis.$fetch = $fetch;
  h3App.use(createRouteRulesHandler({ localFetch }));
  for (const h of handlers) {
    let handler = h.lazy ? lazyEventHandler(h.handler) : h.handler;
    if (h.middleware || !h.route) {
      const middlewareBase = (config.app.baseURL + (h.route || "/")).replace(
        /\/+/g,
        "/"
      );
      h3App.use(middlewareBase, handler);
    } else {
      const routeRules = getRouteRulesForPath(
        h.route.replace(/:\w+|\*\*/g, "_")
      );
      if (routeRules.cache) {
        handler = cachedEventHandler(handler, {
          group: "nitro/routes",
          ...routeRules.cache
        });
      }
      router.use(h.route, handler, h.method);
    }
  }
  h3App.use(config.app.baseURL, router.handler);
  {
    const _handler = h3App.handler;
    h3App.handler = (event) => {
      const ctx = { event };
      return nitroAsyncContext.callAsync(ctx, () => _handler(event));
    };
  }
  const app = {
    hooks,
    h3App,
    router,
    localCall,
    localFetch,
    captureError
  };
  return app;
}
function runNitroPlugins(nitroApp2) {
  for (const plugin of plugins) {
    try {
      plugin(nitroApp2);
    } catch (error) {
      nitroApp2.captureError(error, { tags: ["plugin"] });
      throw error;
    }
  }
}
const nitroApp$1 = createNitroApp();
function useNitroApp() {
  return nitroApp$1;
}
runNitroPlugins(nitroApp$1);

if (!globalThis.crypto) {
  globalThis.crypto = crypto$1;
}
const { NITRO_NO_UNIX_SOCKET, NITRO_DEV_WORKER_ID } = process.env;
trapUnhandledNodeErrors();
parentPort?.on("message", (msg) => {
  if (msg && msg.event === "shutdown") {
    shutdown();
  }
});
const nitroApp = useNitroApp();
const server = new Server(toNodeListener(nitroApp.h3App));
let listener;
listen().catch(() => listen(
  true
  /* use random port */
)).catch((error) => {
  console.error("Dev worker failed to listen:", error);
  return shutdown();
});
nitroApp.router.get(
  "/_nitro/tasks",
  defineEventHandler(async (event) => {
    const _tasks = await Promise.all(
      Object.entries(tasks).map(async ([name, task]) => {
        const _task = await task.resolve?.();
        return [name, { description: _task?.meta?.description }];
      })
    );
    return {
      tasks: Object.fromEntries(_tasks),
      scheduledTasks
    };
  })
);
nitroApp.router.use(
  "/_nitro/tasks/:name",
  defineEventHandler(async (event) => {
    const name = getRouterParam(event, "name");
    const payload = {
      ...getQuery$1(event),
      ...await readBody(event).then((r) => r?.payload).catch(() => ({}))
    };
    return await runTask(name, { payload });
  })
);
function listen(useRandomPort = Boolean(
  NITRO_NO_UNIX_SOCKET || process.versions.webcontainer || "Bun" in globalThis && process.platform === "win32"
)) {
  return new Promise((resolve, reject) => {
    try {
      listener = server.listen(useRandomPort ? 0 : getSocketAddress(), () => {
        const address = server.address();
        parentPort?.postMessage({
          event: "listen",
          address: typeof address === "string" ? { socketPath: address } : { host: "localhost", port: address?.port }
        });
        resolve();
      });
    } catch (error) {
      reject(error);
    }
  });
}
function getSocketAddress() {
  const socketName = `nitro-worker-${process.pid}-${threadId}-${NITRO_DEV_WORKER_ID}-${Math.round(Math.random() * 1e4)}.sock`;
  if (process.platform === "win32") {
    return join(String.raw`\\.\pipe`, socketName);
  }
  if (process.platform === "linux") {
    const nodeMajor = Number.parseInt(process.versions.node.split(".")[0], 10);
    if (nodeMajor >= 20) {
      return `\0${socketName}`;
    }
  }
  return join(tmpdir(), socketName);
}
async function shutdown() {
  server.closeAllConnections?.();
  await Promise.all([
    new Promise((resolve) => listener?.close(resolve)),
    nitroApp.hooks.callHook("close").catch(console.error)
  ]);
  parentPort?.postMessage({ event: "exit" });
}

const _messages = { "appName": "Nuxt", "statusCode": 500, "statusMessage": "Internal server error", "description": "This page is temporarily unavailable.", "refresh": "Refresh this page" };
const template$1 = (messages) => {
  messages = { ..._messages, ...messages };
  return '<!DOCTYPE html><html lang="en"><head><title>' + escapeHtml(messages.statusCode) + " - " + escapeHtml(messages.statusMessage) + " | " + escapeHtml(messages.appName) + `</title><meta charset="utf-8"><meta content="width=device-width,initial-scale=1.0,minimum-scale=1.0" name="viewport"><script>!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))r(e);new MutationObserver(e=>{for(const o of e)if("childList"===o.type)for(const e of o.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&r(e)}).observe(document,{childList:!0,subtree:!0})}function r(e){if(e.ep)return;e.ep=!0;const r=function(e){const r={};return e.integrity&&(r.integrity=e.integrity),e.referrerPolicy&&(r.referrerPolicy=e.referrerPolicy),"use-credentials"===e.crossOrigin?r.credentials="include":"anonymous"===e.crossOrigin?r.credentials="omit":r.credentials="same-origin",r}(e);fetch(e.href,r)}}();<\/script><style>*,:after,:before{border-color:var(--un-default-border-color,#e5e7eb);border-style:solid;border-width:0;box-sizing:border-box}:after,:before{--un-content:""}html{line-height:1.5;-webkit-text-size-adjust:100%;font-family:ui-sans-serif,system-ui,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol,Noto Color Emoji;font-feature-settings:normal;font-variation-settings:normal;-moz-tab-size:4;tab-size:4;-webkit-tap-highlight-color:transparent}body{line-height:inherit;margin:0}h1,h2{font-size:inherit;font-weight:inherit}h1,h2,p{margin:0}*,:after,:before{--un-rotate:0;--un-rotate-x:0;--un-rotate-y:0;--un-rotate-z:0;--un-scale-x:1;--un-scale-y:1;--un-scale-z:1;--un-skew-x:0;--un-skew-y:0;--un-translate-x:0;--un-translate-y:0;--un-translate-z:0;--un-pan-x: ;--un-pan-y: ;--un-pinch-zoom: ;--un-scroll-snap-strictness:proximity;--un-ordinal: ;--un-slashed-zero: ;--un-numeric-figure: ;--un-numeric-spacing: ;--un-numeric-fraction: ;--un-border-spacing-x:0;--un-border-spacing-y:0;--un-ring-offset-shadow:0 0 transparent;--un-ring-shadow:0 0 transparent;--un-shadow-inset: ;--un-shadow:0 0 transparent;--un-ring-inset: ;--un-ring-offset-width:0px;--un-ring-offset-color:#fff;--un-ring-width:0px;--un-ring-color:rgba(147,197,253,.5);--un-blur: ;--un-brightness: ;--un-contrast: ;--un-drop-shadow: ;--un-grayscale: ;--un-hue-rotate: ;--un-invert: ;--un-saturate: ;--un-sepia: ;--un-backdrop-blur: ;--un-backdrop-brightness: ;--un-backdrop-contrast: ;--un-backdrop-grayscale: ;--un-backdrop-hue-rotate: ;--un-backdrop-invert: ;--un-backdrop-opacity: ;--un-backdrop-saturate: ;--un-backdrop-sepia: }.grid{display:grid}.mb-2{margin-bottom:.5rem}.mb-4{margin-bottom:1rem}.max-w-520px{max-width:520px}.min-h-screen{min-height:100vh}.place-content-center{place-content:center}.overflow-hidden{overflow:hidden}.bg-white{--un-bg-opacity:1;background-color:rgb(255 255 255/var(--un-bg-opacity))}.px-2{padding-left:.5rem;padding-right:.5rem}.text-center{text-align:center}.text-\\[80px\\]{font-size:80px}.text-2xl{font-size:1.5rem;line-height:2rem}.text-\\[\\#020420\\]{--un-text-opacity:1;color:rgb(2 4 32/var(--un-text-opacity))}.text-\\[\\#64748B\\]{--un-text-opacity:1;color:rgb(100 116 139/var(--un-text-opacity))}.font-semibold{font-weight:600}.leading-none{line-height:1}.tracking-wide{letter-spacing:.025em}.font-sans{font-family:ui-sans-serif,system-ui,-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Helvetica Neue,Arial,Noto Sans,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol,Noto Color Emoji}.tabular-nums{--un-numeric-spacing:tabular-nums;font-variant-numeric:var(--un-ordinal) var(--un-slashed-zero) var(--un-numeric-figure) var(--un-numeric-spacing) var(--un-numeric-fraction)}.antialiased{-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale}@media(prefers-color-scheme:dark){.dark\\:bg-\\[\\#020420\\]{--un-bg-opacity:1;background-color:rgb(2 4 32/var(--un-bg-opacity))}.dark\\:text-white{--un-text-opacity:1;color:rgb(255 255 255/var(--un-text-opacity))}}@media(min-width:640px){.sm\\:text-\\[110px\\]{font-size:110px}.sm\\:text-3xl{font-size:1.875rem;line-height:2.25rem}}</style></head><body class="antialiased bg-white dark:bg-[#020420] dark:text-white font-sans grid min-h-screen overflow-hidden place-content-center text-[#020420] tracking-wide"><div class="max-w-520px text-center"><h1 class="font-semibold leading-none mb-4 sm:text-[110px] tabular-nums text-[80px]">` + escapeHtml(messages.statusCode) + '</h1><h2 class="font-semibold mb-2 sm:text-3xl text-2xl">' + escapeHtml(messages.statusMessage) + '</h2><p class="mb-4 px-2 text-[#64748B] text-md">' + escapeHtml(messages.description) + "</p></div></body></html>";
};

const error500 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  template: template$1
}, Symbol.toStringTag, { value: 'Module' }));

const template = "";

const _virtual__spaTemplate = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  template: template
}, Symbol.toStringTag, { value: 'Module' }));

const styles = {};

const styles$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: styles
}, Symbol.toStringTag, { value: 'Module' }));

const login_post = defineEventHandler(async (event) => {
  const body = await readBody(event);
  if (!body.username || !body.password) {
    throw createError({
      statusCode: 400,
      message: "\u7528\u6237\u540D\u548C\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A"
    });
  }
  try {
    const res = await apiRequest(event, "/gateway/user/v1/login", {
      method: "POST",
      body: {
        username: body.username,
        password: body.password
      }
    });
    const { token, user } = res.data;
    await setUserSession(event, {
      user: {
        provider: "local",
        id: String(user.id),
        name: user.nickName || user.userName,
        userName: user.userName,
        nickName: user.nickName,
        avatar: user.avatar,
        email: user.email,
        signature: user.signature,
        backgroundImage: user.backgroundImage,
        phone: user.phone
      },
      token
    });
    return {
      code: 200,
      data: {
        token,
        user
      },
      msg: "\u767B\u5F55\u6210\u529F"
    };
  } catch (error) {
    throw createError({
      statusCode: 401,
      message: error.message || "\u767B\u5F55\u5931\u8D25"
    });
  }
});

const login_post$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: login_post
}, Symbol.toStringTag, { value: 'Module' }));

const register_post = defineEventHandler(async (event) => {
  const body = await readBody(event);
  if (!body.username || !body.password) {
    throw createError({
      statusCode: 400,
      message: "\u7528\u6237\u540D\u548C\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A"
    });
  }
  try {
    const res = await apiRequest(event, "/gateway/user/v1/register", {
      method: "POST",
      body: {
        username: body.username,
        password: body.password
      }
    });
    const { token, user } = res.data;
    await setUserSession(event, {
      user: {
        provider: "local",
        id: String(user.id),
        name: user.nickName || user.userName,
        userName: user.userName,
        nickName: user.nickName,
        avatar: user.avatar,
        email: user.email,
        signature: user.signature,
        backgroundImage: user.backgroundImage,
        phone: user.phone
      },
      token
    });
    return {
      code: 200,
      data: {
        token,
        user
      },
      msg: "\u6CE8\u518C\u6210\u529F"
    };
  } catch (error) {
    throw createError({
      statusCode: 400,
      message: error.message || "\u6CE8\u518C\u5931\u8D25"
    });
  }
});

const register_post$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: register_post
}, Symbol.toStringTag, { value: 'Module' }));

const drawings_get = eventHandler(async (event) => {
  const { cursor } = await getQuery$1(event);
  const res = await blob.list({
    limit: 20,
    cursor,
    prefix: "drawings/"
  });
  res.blobs = await Promise.all(res.blobs.map((b) => !Object.keys(b.customMetadata).length ? blob.head(b.pathname) : b));
  res.blobs = res.blobs.map((b) => {
    var _a, _b, _c, _d, _e, _f, _g, _h, _i, _j, _k, _l, _m, _n, _o, _p, _q, _r, _s, _t;
    const aiImage = ((_a = b.customMetadata) == null ? void 0 : _a.aiImage) || ((_b = b.customMetadata) == null ? void 0 : _b.aiimage) || ((_c = b.customMetadata) == null ? void 0 : _c.ai_image);
    return {
      ...b,
      customMetadata: {
        url: process.env.S3_PUBLIC_URL ? `${process.env.S3_PUBLIC_URL}/${b.pathname}` : void 0,
        userProvider: ((_d = b.customMetadata) == null ? void 0 : _d.userProvider) || ((_e = b.customMetadata) == null ? void 0 : _e.userprovider) || ((_f = b.customMetadata) == null ? void 0 : _f.user_provider),
        userId: ((_g = b.customMetadata) == null ? void 0 : _g.userId) || ((_h = b.customMetadata) == null ? void 0 : _h.userid) || ((_i = b.customMetadata) == null ? void 0 : _i.user_id),
        userName: ((_j = b.customMetadata) == null ? void 0 : _j.userName) || ((_k = b.customMetadata) == null ? void 0 : _k.username) || ((_l = b.customMetadata) == null ? void 0 : _l.user_name),
        userAvatar: ((_m = b.customMetadata) == null ? void 0 : _m.userAvatar) || ((_n = b.customMetadata) == null ? void 0 : _n.useravatar) || ((_o = b.customMetadata) == null ? void 0 : _o.user_avatar),
        userUrl: ((_p = b.customMetadata) == null ? void 0 : _p.userUrl) || ((_q = b.customMetadata) == null ? void 0 : _q.userurl) || ((_r = b.customMetadata) == null ? void 0 : _r.user_url),
        description: (_s = b.customMetadata) == null ? void 0 : _s.description,
        aiImage,
        aiImageUrl: ((_t = b.customMetadata) == null ? void 0 : _t.aiImageUrl) || (process.env.S3_PUBLIC_URL ? `${process.env.S3_PUBLIC_URL}/${aiImage}` : void 0)
      }
    };
  });
  return res;
});

const drawings_get$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: drawings_get
}, Symbol.toStringTag, { value: 'Module' }));

const feed_get = defineEventHandler(async (event) => {
  var _a;
  const query = await getQuery$1(event);
  const latestTime = query.latestTime ? Number.parseInt(query.latestTime) : Date.now();
  const page = query.page ? Number.parseInt(query.page) : 1;
  const size = query.size ? Number.parseInt(query.size) : 10;
  try {
    const queryParams = new URLSearchParams({
      latestTime: latestTime.toString(),
      page: page.toString(),
      size: size.toString()
    });
    const res = await apiRequest(
      event,
      `/gateway/video/v1/feed?${queryParams.toString()}`,
      {
        method: "GET",
        requireAuth: false
        // 不需要认证，但如果有token会自动携带
      }
    );
    return {
      code: 200,
      data: res.data,
      msg: "\u67E5\u8BE2\u6210\u529F"
    };
  } catch (error) {
    console.error("[Feed API Error]", {
      message: error.message,
      statusCode: error.statusCode,
      data: error.data,
      stack: error.stack
    });
    const statusCode = error.statusCode || (((_a = error.message) == null ? void 0 : _a.includes("\u9700\u8981\u767B\u5F55")) ? 401 : 500);
    const message = error.message || "\u83B7\u53D6feed\u5931\u8D25";
    throw createError({
      statusCode,
      message
    });
  }
});

const feed_get$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: feed_get
}, Symbol.toStringTag, { value: 'Module' }));

const generate_post = eventHandler(async (event) => {
  await requireUserSession(event);
  const form = await readFormData(event);
  const drawing = form.get("drawing");
  const drawingArrayBuffer = await drawing.arrayBuffer();
  const { text } = await generateText({
    model: hubAI("openai/gpt-5-nano"),
    prompt: [{
      role: "user",
      content: "Describe this drawing in one sentence."
    }, {
      role: "user",
      content: [{
        type: "image",
        image: drawingArrayBuffer
      }]
    }]
  }).catch(() => {
    return { text: "" };
  });
  setHeader(event, "x-description", text);
  const result = await generateText({
    model: hubAI("google/gemini-2.5-flash-image-preview"),
    providerOptions: {
      google: { responseModalities: ["TEXT", "IMAGE"] }
    },
    messages: [
      {
        role: "user",
        content: [
          {
            type: "text",
            text: "Generate a more detailed and beautiful image of the following drawing:"
          },
          {
            type: "image",
            image: drawingArrayBuffer
          }
        ]
      }
    ]
  });
  const image = result.files[0];
  if (!image) {
    throw createError({
      statusCode: 500,
      message: "Failed to generate image"
    });
  }
  setHeader(event, "content-type", image.mediaType);
  return image.uint8Array;
});

const generate_post$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: generate_post
}, Symbol.toStringTag, { value: 'Module' }));

const upload_post = eventHandler(async (event) => {
  var _b;
  const { user } = await requireUserSession(event);
  const { blobs } = await blob.list({
    prefix: "drawings/",
    limit: 1
  });
  const form = await readFormData(event);
  const drawing = form.get("drawing");
  const drawingArrayBuffer = await drawing.arrayBuffer();
  ensureBlob(drawing, {
    maxSize: "1MB",
    types: ["image/jpeg"]
  });
  const { text } = await generateText({
    model: "openai/gpt-5-nano",
    prompt: [{
      role: "user",
      content: "Describe this drawing in one sentence."
    }, {
      role: "user",
      content: [{
        type: "image",
        image: drawingArrayBuffer
      }]
    }]
  }).catch(() => {
    return { text: "" };
  });
  if (text.includes("penis")) {
    throw createError({
      statusCode: 400,
      message: "You cannot upload this kind of drawings."
    });
  }
  const name = `${(/* @__PURE__ */ new Date("2050-01-01")).getTime() - Date.now()}`;
  const aiImage = await generateText({
    model: "google/gemini-3-pro-image",
    providerOptions: {
      google: { responseModalities: ["TEXT", "IMAGE"] }
    },
    messages: [
      {
        role: "user",
        content: [
          {
            type: "text",
            text: "Generate a more detailed and beautiful drawing based on the following drawing. Try to respect the drawing as much as possible. Please generate the image in jpg, a 1:1 aspect ratio, as light as possible in term of size and 512x512px."
          },
          {
            type: "image",
            image: drawingArrayBuffer
          }
        ]
      }
    ]
  }).then((result) => {
    const generatedImage = result.files[0];
    return blob.put(`${name}.${generatedImage.mediaType.split("/")[1]}`, generatedImage.uint8Array, {
      prefix: "ai/",
      addRandomSuffix: true,
      contentType: generatedImage.mediaType
    });
  }).catch(() => null);
  console.log("blob put", `${name}.jpg`, drawing);
  return blob.put(`${name}.jpg`, drawing, {
    prefix: "drawings/",
    addRandomSuffix: true,
    customMetadata: {
      userProvider: user.provider,
      userId: user.id,
      userName: user.name,
      userAvatar: user.avatar,
      userUrl: user.url,
      description: text.trim(),
      aiImage: aiImage ? aiImage.pathname : "",
      aiImageUrl: ((_b = aiImage == null ? void 0 : aiImage.customMetadata) == null ? void 0 : _b.url) || ""
    }
  });
});

const upload_post$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: upload_post
}, Symbol.toStringTag, { value: 'Module' }));

const anonymous_get = eventHandler(async (event) => {
  var _a, _b;
  useRuntimeConfig();
  const fingerprint = await getRequestFingerprint(event, {
    userAgent: true,
    xForwardedFor: true
  });
  await setUserSession(event, {
    user: {
      provider: "anonymous",
      id: fingerprint || "anonymous",
      name: `${(_a = event.context.cf) == null ? void 0 : _a.city}, ${(_b = event.context.cf) == null ? void 0 : _b.country}`,
      avatar: "",
      url: ""
    }
  });
  return sendRedirect(event, "/draw");
});

const anonymous_get$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: anonymous_get
}, Symbol.toStringTag, { value: 'Module' }));

const github_get = defineOAuthGitHubEventHandler({
  async onSuccess(event, { user }) {
    await setUserSession(event, {
      user: {
        provider: "github",
        id: String(user.id),
        name: user.name || user.login,
        avatar: user.avatar_url,
        url: user.html_url
      }
    });
    return sendRedirect(event, "/draw");
  }
});

const github_get$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: github_get
}, Symbol.toStringTag, { value: 'Module' }));

const google_get = defineOAuthGoogleEventHandler({
  async onSuccess(event, { user }) {
    await setUserSession(event, {
      user: {
        provider: "google",
        id: user.sub,
        name: user.name,
        avatar: user.picture,
        url: ""
      }
    });
    return sendRedirect(event, "/draw");
  }
});

const google_get$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: google_get
}, Symbol.toStringTag, { value: 'Module' }));

const ____pathname__get = eventHandler(async (event) => {
  const { pathname } = event.context.params || {};
  return blob.serve(event, pathname);
});

const ____pathname__get$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: ____pathname__get
}, Symbol.toStringTag, { value: 'Module' }));

function renderPayloadResponse(ssrContext) {
  return {
    body: stringify(splitPayload(ssrContext).payload, ssrContext._payloadReducers) ,
    statusCode: getResponseStatus(ssrContext.event),
    statusMessage: getResponseStatusText(ssrContext.event),
    headers: {
      "content-type": "application/json;charset=utf-8" ,
      "x-powered-by": "Nuxt"
    }
  };
}
function renderPayloadJsonScript(opts) {
  const contents = opts.data ? stringify(opts.data, opts.ssrContext._payloadReducers) : "";
  const payload = {
    "type": "application/json",
    "innerHTML": contents,
    "data-nuxt-data": appId,
    "data-ssr": !(opts.ssrContext.noSSR)
  };
  {
    payload.id = "__NUXT_DATA__";
  }
  if (opts.src) {
    payload["data-src"] = opts.src;
  }
  const config = uneval(opts.ssrContext.config);
  return [
    payload,
    {
      innerHTML: `window.__NUXT__={};window.__NUXT__.config=${config}`
    }
  ];
}
function splitPayload(ssrContext) {
  const { data, prerenderedAt, ...initial } = ssrContext.payload;
  return {
    initial: { ...initial, prerenderedAt },
    payload: { data, prerenderedAt }
  };
}

const renderSSRHeadOptions = {"omitLineBreaks":true};

globalThis.__buildAssetsURL = buildAssetsURL;
globalThis.__publicAssetsURL = publicAssetsURL;
const HAS_APP_TELEPORTS = !!(appTeleportAttrs.id);
const APP_TELEPORT_OPEN_TAG = HAS_APP_TELEPORTS ? `<${appTeleportTag}${propsToString(appTeleportAttrs)}>` : "";
const APP_TELEPORT_CLOSE_TAG = HAS_APP_TELEPORTS ? `</${appTeleportTag}>` : "";
const PAYLOAD_URL_RE = /^[^?]*\/_payload.json(?:\?.*)?$/ ;
const renderer = defineRenderHandler(async (event) => {
  const nitroApp = useNitroApp();
  const ssrError = event.path.startsWith("/__nuxt_error") ? getQuery$1(event) : null;
  if (ssrError && !("__unenv__" in event.node.req)) {
    throw createError({
      statusCode: 404,
      statusMessage: "Page Not Found: /__nuxt_error"
    });
  }
  const ssrContext = createSSRContext(event);
  const headEntryOptions = { mode: "server" };
  ssrContext.head.push(appHead, headEntryOptions);
  if (ssrError) {
    ssrError.statusCode &&= Number.parseInt(ssrError.statusCode);
    if (typeof ssrError.data === "string") {
      try {
        ssrError.data = destr(ssrError.data);
      } catch {
      }
    }
    setSSRError(ssrContext, ssrError);
  }
  const isRenderingPayload = PAYLOAD_URL_RE.test(ssrContext.url);
  if (isRenderingPayload) {
    const url = ssrContext.url.substring(0, ssrContext.url.lastIndexOf("/")) || "/";
    ssrContext.url = url;
    event._path = event.node.req.url = url;
  }
  const routeOptions = getRouteRules(event);
  if (routeOptions.ssr === false) {
    ssrContext.noSSR = true;
  }
  const renderer = await getRenderer(ssrContext);
  const _rendered = await renderer.renderToString(ssrContext).catch(async (error) => {
    if (ssrContext._renderResponse && error.message === "skipping render") {
      return {};
    }
    const _err = !ssrError && ssrContext.payload?.error || error;
    await ssrContext.nuxt?.hooks.callHook("app:error", _err);
    throw _err;
  });
  const inlinedStyles = [];
  await ssrContext.nuxt?.hooks.callHook("app:rendered", { ssrContext, renderResult: _rendered });
  if (ssrContext._renderResponse) {
    return ssrContext._renderResponse;
  }
  if (ssrContext.payload?.error && !ssrError) {
    throw ssrContext.payload.error;
  }
  if (isRenderingPayload) {
    const response = renderPayloadResponse(ssrContext);
    return response;
  }
  const NO_SCRIPTS = routeOptions.noScripts;
  const { styles, scripts } = getRequestDependencies(ssrContext, renderer.rendererContext);
  if (ssrContext._preloadManifest && !NO_SCRIPTS) {
    ssrContext.head.push({
      link: [
        { rel: "preload", as: "fetch", fetchpriority: "low", crossorigin: "anonymous", href: buildAssetsURL(`builds/meta/${ssrContext.runtimeConfig.app.buildId}.json`) }
      ]
    }, { ...headEntryOptions, tagPriority: "low" });
  }
  if (inlinedStyles.length) {
    ssrContext.head.push({ style: inlinedStyles });
  }
  const link = [];
  for (const resource of Object.values(styles)) {
    if ("inline" in getQuery(resource.file)) {
      continue;
    }
    link.push({ rel: "stylesheet", href: renderer.rendererContext.buildAssetsURL(resource.file), crossorigin: "" });
  }
  if (link.length) {
    ssrContext.head.push({ link }, headEntryOptions);
  }
  if (!NO_SCRIPTS) {
    ssrContext.head.push({
      link: getPreloadLinks(ssrContext, renderer.rendererContext)
    }, headEntryOptions);
    ssrContext.head.push({
      link: getPrefetchLinks(ssrContext, renderer.rendererContext)
    }, headEntryOptions);
    ssrContext.head.push({
      script: renderPayloadJsonScript({ ssrContext, data: ssrContext.payload }) 
    }, {
      ...headEntryOptions,
      // this should come before another end of body scripts
      tagPosition: "bodyClose",
      tagPriority: "high"
    });
  }
  if (!routeOptions.noScripts) {
    const tagPosition = "head";
    ssrContext.head.push({
      script: Object.values(scripts).map((resource) => ({
        type: resource.module ? "module" : null,
        src: renderer.rendererContext.buildAssetsURL(resource.file),
        defer: resource.module ? null : true,
        // if we are rendering script tag payloads that import an async payload
        // we need to ensure this resolves before executing the Nuxt entry
        tagPosition,
        crossorigin: ""
      }))
    }, headEntryOptions);
  }
  const { headTags, bodyTags, bodyTagsOpen, htmlAttrs, bodyAttrs } = await renderSSRHead(ssrContext.head, renderSSRHeadOptions);
  const htmlContext = {
    htmlAttrs: htmlAttrs ? [htmlAttrs] : [],
    head: normalizeChunks([headTags]),
    bodyAttrs: bodyAttrs ? [bodyAttrs] : [],
    bodyPrepend: normalizeChunks([bodyTagsOpen, ssrContext.teleports?.body]),
    body: [
      replaceIslandTeleports(ssrContext, _rendered.html) ,
      APP_TELEPORT_OPEN_TAG + (HAS_APP_TELEPORTS ? joinTags([ssrContext.teleports?.[`#${appTeleportAttrs.id}`]]) : "") + APP_TELEPORT_CLOSE_TAG
    ],
    bodyAppend: [bodyTags]
  };
  await nitroApp.hooks.callHook("render:html", htmlContext, { event });
  return {
    body: renderHTMLDocument(htmlContext),
    statusCode: getResponseStatus(event),
    statusMessage: getResponseStatusText(event),
    headers: {
      "content-type": "text/html;charset=utf-8",
      "x-powered-by": "Nuxt"
    }
  };
});
function normalizeChunks(chunks) {
  const result = [];
  for (const _chunk of chunks) {
    const chunk = _chunk?.trim();
    if (chunk) {
      result.push(chunk);
    }
  }
  return result;
}
function joinTags(tags) {
  return tags.join("");
}
function joinAttrs(chunks) {
  if (chunks.length === 0) {
    return "";
  }
  return " " + chunks.join(" ");
}
function renderHTMLDocument(html) {
  return `<!DOCTYPE html><html${joinAttrs(html.htmlAttrs)}><head>${joinTags(html.head)}</head><body${joinAttrs(html.bodyAttrs)}>${joinTags(html.bodyPrepend)}${joinTags(html.body)}${joinTags(html.bodyAppend)}</body></html>`;
}

const renderer$1 = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: renderer
}, Symbol.toStringTag, { value: 'Module' }));
//# sourceMappingURL=index.mjs.map
