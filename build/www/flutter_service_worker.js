'use strict';
const MANIFEST = 'flutter-app-manifest';
const TEMP = 'flutter-temp-cache';
const CACHE_NAME = 'flutter-app-cache';

const RESOURCES = {"assets/AssetManifest.bin": "1121a4b9dee63498de9a739f84783849",
"assets/AssetManifest.json": "b3c480800f81358d83d94b4c483d4352",
"assets/FontManifest.json": "fcc804d5e4c25e196e698d253a79cbd8",
"assets/fonts/MaterialIcons-Regular.otf": "84bf4b09036d502e04d02b8418c6b800",
"assets/NOTICES": "dfcee7e1d0a1c15ca5be204d241612a0",
"assets/packages/font_awesome_flutter/lib/fonts/fa-brands-400.ttf": "8c0c780f82aa8aaa2dc3f638506ea952",
"assets/packages/font_awesome_flutter/lib/fonts/fa-regular-400.ttf": "b7a50a5b33477fb1ceea0cad4bead825",
"assets/packages/font_awesome_flutter/lib/fonts/fa-solid-900.ttf": "0eae886b7c45b71b33a548584a9214b7",
"assets/resources/app_icons/chrome.png": "02b40e735be29f8b5729505278f1d5ee",
"assets/resources/app_icons/excel.png": "f1de1dcfb109ffde7e90a80d2c3a7853",
"assets/resources/fonts/Comic-L.ttf": "a50f9c96a76356e3d01013e0b042989f",
"assets/resources/fonts/Roboto-Black.ttf": "d6a6f8878adb0d8e69f9fa2e0b622924",
"assets/resources/fonts/Roboto-BlackItalic.ttf": "c3332e3b8feff748ecb0c6cb75d65eae",
"assets/resources/fonts/Roboto-Bold.ttf": "b8e42971dec8d49207a8c8e2b919a6ac",
"assets/resources/fonts/Roboto-BoldItalic.ttf": "fd6e9700781c4aaae877999d09db9e09",
"assets/resources/fonts/Roboto-Italic.ttf": "cebd892d1acfcc455f5e52d4104f2719",
"assets/resources/fonts/Roboto-Light.ttf": "881e150ab929e26d1f812c4342c15a7c",
"assets/resources/fonts/Roboto-LightItalic.ttf": "5788d5ce921d7a9b4fa0eaa9bf7fec8d",
"assets/resources/fonts/Roboto-Medium.ttf": "68ea4734cf86bd544650aee05137d7bb",
"assets/resources/fonts/Roboto-MediumItalic.ttf": "c16d19c2c0fd1278390a82fc245f4923",
"assets/resources/fonts/Roboto-Regular.ttf": "8a36205bd9b83e03af0591a004bc97f4",
"assets/resources/fonts/Roboto-Thin.ttf": "66209ae01f484e46679622dd607fcbc5",
"assets/resources/fonts/Roboto-ThinItalic.ttf": "7bcadd0675fe47d69c2d8aaef683416f",
"assets/resources/fonts/Ubuntu-B.ttf": "008e6bc48c8eaa5d2855d57e6b0b8595",
"assets/resources/fonts/Ubuntu-L.ttf": "2759de5c01527bd9730b4d1838e6c938",
"assets/resources/i18n/en.json": "6a15667303f6d07ae808e01ab0f2b71a",
"assets/resources/i18n/pt.json": "98471e5af3802ae5dbd53b5edc6ac506",
"assets/resources/icons/desktop.png": "2d215165803d1157655d644376f8b171",
"assets/resources/icons/laptop.png": "10c64c92f464cb056acaa4f4dab3c5ae",
"assets/resources/icons/user_idle.png": "33772db48589a605a1a7770a6353a066",
"assets/resources/icons/user_idle_solid.png": "d66b138f5a3d2512b2c14c7b90970927",
"assets/resources/images/atendas_logo.png": "73eaf6e1efc05e86e35327aedebe19cf",
"assets/resources/images/atendas_slider.png": "b2d9e47e1f50e237f785e9abd6159c8c",
"assets/resources/images/atendas_text_logo.png": "8be27d89fa3f2feb32832adcfe7d08ae",
"assets/resources/images/atendas_text_logo_dark.png": "6dd56fb13ea97f2720f941e479927f12",
"assets/resources/images/makrolock_logo.png": "50d3476d15747299d79094fae54808fe",
"assets/resources/images/makrolock_slider.png": "4f0de4c09e5e34fe29277f5983e4e972",
"assets/resources/images/makrolock_text_logo.png": "0540c845ca5b100d04da90729e4a6151",
"assets/resources/images/makrolock_text_logo_alternative.png": "d3d4c9a7cef62f1d24eb68440e8549ee",
"assets/resources/images/makrolock_text_logo_dark.png": "e1e6af679da16cf29c099e89a8be9934",
"assets/resources/images/offline.png": "ccf23270d1b427efe71bf59f444304b7",
"assets/resources/images/service_logo.png": "8ea53e612c36654c977b94b0b4b63880",
"assets/resources/images/service_slider.png": "a2089a347c5a5708ae8969e72ac61b8a",
"assets/resources/images/service_text_logo.png": "b14c978b2149c5cb5dda67f81610dbdf",
"assets/resources/images/service_text_logo_dark.png": "b14c978b2149c5cb5dda67f81610dbdf",
"assets/shaders/ink_sparkle.frag": "f8b80e740d33eb157090be4e995febdf",
"canvaskit/canvaskit.js": "bbf39143dfd758d8d847453b120c8ebb",
"canvaskit/canvaskit.wasm": "42df12e09ecc0d5a4a34a69d7ee44314",
"canvaskit/chromium/canvaskit.js": "96ae916cd2d1b7320fff853ee22aebb0",
"canvaskit/chromium/canvaskit.wasm": "be0e3b33510f5b7b0cc76cc4d3e50048",
"canvaskit/skwasm.js": "95f16c6690f955a45b2317496983dbe9",
"canvaskit/skwasm.wasm": "1a074e8452fe5e0d02b112e22cdcf455",
"canvaskit/skwasm.worker.js": "51253d3321b11ddb8d73fa8aa87d3b15",
"favicon.ico": "c50c0a10322410d13791861475800ec2",
"favicon.png": "31615eff466174b021599f42b50aa278",
"favicon_atendas.ico": "64d5cf09c6522481687ad156ac88eda4",
"favicon_service.ico": "cde0ce1aac3ddd5ffadd8bd647726d51",
"flutter.js": "6b515e434cea20006b3ef1726d2c8894",
"icons/Icon-192.png": "2a0873577f664acdda7bd6dde5ec7024",
"icons/Icon-512.png": "1b4d16662785bcc9f52ad64f7781ab32",
"icons/Icon-maskable-192.png": "c457ef57daa1d16f64b27b786ec2ea3c",
"icons/Icon-maskable-512.png": "301a7604d45b3e739efc881eb04896ea",
"index.html": "589d6d276315d8001df6f33af268d73f",
"/": "589d6d276315d8001df6f33af268d73f",
"main.dart.js": "6e8db927db1fac68850ab85febc1163a",
"manifest.json": "328cd33789873777738998ad0d623c61",
"version.json": "3901f42d065f6f95d75cc2ba3522fc10"};
// The application shell files that are downloaded before a service worker can
// start.
const CORE = ["main.dart.js",
"index.html",
"assets/AssetManifest.json",
"assets/FontManifest.json"];

// During install, the TEMP cache is populated with the application shell files.
self.addEventListener("install", (event) => {
  self.skipWaiting();
  return event.waitUntil(
    caches.open(TEMP).then((cache) => {
      return cache.addAll(
        CORE.map((value) => new Request(value, {'cache': 'reload'})));
    })
  );
});
// During activate, the cache is populated with the temp files downloaded in
// install. If this service worker is upgrading from one with a saved
// MANIFEST, then use this to retain unchanged resource files.
self.addEventListener("activate", function(event) {
  return event.waitUntil(async function() {
    try {
      var contentCache = await caches.open(CACHE_NAME);
      var tempCache = await caches.open(TEMP);
      var manifestCache = await caches.open(MANIFEST);
      var manifest = await manifestCache.match('manifest');
      // When there is no prior manifest, clear the entire cache.
      if (!manifest) {
        await caches.delete(CACHE_NAME);
        contentCache = await caches.open(CACHE_NAME);
        for (var request of await tempCache.keys()) {
          var response = await tempCache.match(request);
          await contentCache.put(request, response);
        }
        await caches.delete(TEMP);
        // Save the manifest to make future upgrades efficient.
        await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
        // Claim client to enable caching on first launch
        self.clients.claim();
        return;
      }
      var oldManifest = await manifest.json();
      var origin = self.location.origin;
      for (var request of await contentCache.keys()) {
        var key = request.url.substring(origin.length + 1);
        if (key == "") {
          key = "/";
        }
        // If a resource from the old manifest is not in the new cache, or if
        // the MD5 sum has changed, delete it. Otherwise the resource is left
        // in the cache and can be reused by the new service worker.
        if (!RESOURCES[key] || RESOURCES[key] != oldManifest[key]) {
          await contentCache.delete(request);
        }
      }
      // Populate the cache with the app shell TEMP files, potentially overwriting
      // cache files preserved above.
      for (var request of await tempCache.keys()) {
        var response = await tempCache.match(request);
        await contentCache.put(request, response);
      }
      await caches.delete(TEMP);
      // Save the manifest to make future upgrades efficient.
      await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
      // Claim client to enable caching on first launch
      self.clients.claim();
      return;
    } catch (err) {
      // On an unhandled exception the state of the cache cannot be guaranteed.
      console.error('Failed to upgrade service worker: ' + err);
      await caches.delete(CACHE_NAME);
      await caches.delete(TEMP);
      await caches.delete(MANIFEST);
    }
  }());
});
// The fetch handler redirects requests for RESOURCE files to the service
// worker cache.
self.addEventListener("fetch", (event) => {
  if (event.request.method !== 'GET') {
    return;
  }
  var origin = self.location.origin;
  var key = event.request.url.substring(origin.length + 1);
  // Redirect URLs to the index.html
  if (key.indexOf('?v=') != -1) {
    key = key.split('?v=')[0];
  }
  if (event.request.url == origin || event.request.url.startsWith(origin + '/#') || key == '') {
    key = '/';
  }
  // If the URL is not the RESOURCE list then return to signal that the
  // browser should take over.
  if (!RESOURCES[key]) {
    return;
  }
  // If the URL is the index.html, perform an online-first request.
  if (key == '/') {
    return onlineFirst(event);
  }
  event.respondWith(caches.open(CACHE_NAME)
    .then((cache) =>  {
      return cache.match(event.request).then((response) => {
        // Either respond with the cached resource, or perform a fetch and
        // lazily populate the cache only if the resource was successfully fetched.
        return response || fetch(event.request).then((response) => {
          if (response && Boolean(response.ok)) {
            cache.put(event.request, response.clone());
          }
          return response;
        });
      })
    })
  );
});
self.addEventListener('message', (event) => {
  // SkipWaiting can be used to immediately activate a waiting service worker.
  // This will also require a page refresh triggered by the main worker.
  if (event.data === 'skipWaiting') {
    self.skipWaiting();
    return;
  }
  if (event.data === 'downloadOffline') {
    downloadOffline();
    return;
  }
});
// Download offline will check the RESOURCES for all files not in the cache
// and populate them.
async function downloadOffline() {
  var resources = [];
  var contentCache = await caches.open(CACHE_NAME);
  var currentContent = {};
  for (var request of await contentCache.keys()) {
    var key = request.url.substring(origin.length + 1);
    if (key == "") {
      key = "/";
    }
    currentContent[key] = true;
  }
  for (var resourceKey of Object.keys(RESOURCES)) {
    if (!currentContent[resourceKey]) {
      resources.push(resourceKey);
    }
  }
  return contentCache.addAll(resources);
}
// Attempt to download the resource online before falling back to
// the offline cache.
function onlineFirst(event) {
  return event.respondWith(
    fetch(event.request).then((response) => {
      return caches.open(CACHE_NAME).then((cache) => {
        cache.put(event.request, response.clone());
        return response;
      });
    }).catch((error) => {
      return caches.open(CACHE_NAME).then((cache) => {
        return cache.match(event.request).then((response) => {
          if (response != null) {
            return response;
          }
          throw error;
        });
      });
    })
  );
}
