# Changelog

## 1.0.0 (2023-12-04)


### ⚠ BREAKING CHANGES

* OFO APIs were updated to version v1beta1, since they are more stable now. Resources of the alpha versions are no longer supported in flagd or flagd-proxy.
* rename metrics and service ([#730](https://github.com/Kavindu-Dodan/flagd/issues/730))
* proposal - rename `kube-flagd-proxy` to `flagd-proxy` ([#576](https://github.com/Kavindu-Dodan/flagd/issues/576))
* unify sources configuration handling ([#560](https://github.com/Kavindu-Dodan/flagd/issues/560))
* Use OTel to export metrics (metric name changes) ([#419](https://github.com/Kavindu-Dodan/flagd/issues/419))

### Features

* add `$flagd.timestamp` to json evaluator ([#958](https://github.com/Kavindu-Dodan/flagd/issues/958)) ([a1b04e7](https://github.com/Kavindu-Dodan/flagd/commit/a1b04e778df4d2f6beb881183a82161018151479))
* add `sem_ver` jsonLogic evaluator ([#675](https://github.com/Kavindu-Dodan/flagd/issues/675)) ([a8d8ab6](https://github.com/Kavindu-Dodan/flagd/commit/a8d8ab6b4495457a40a2c32b8bd5be48b1fd6941))
* add `starts_with` and `ends_with` json evaluators ([#658](https://github.com/Kavindu-Dodan/flagd/issues/658)) ([f932b8f](https://github.com/Kavindu-Dodan/flagd/commit/f932b8f4c834a5ebe27ebb860c26fdea8da20598))
* add debug logging for merge behaviour  ([#456](https://github.com/Kavindu-Dodan/flagd/issues/456)) ([dc71e84](https://github.com/Kavindu-Dodan/flagd/commit/dc71e84f0704690b528e7f1c2b56cb4898374fbf))
* add flag key to hash in fractional evaluation ([#847](https://github.com/Kavindu-Dodan/flagd/issues/847)) ([ca6a35f](https://github.com/Kavindu-Dodan/flagd/commit/ca6a35fd72462177f45a116e9009fc30b3588b83))
* add gRPC healthchecks ([#863](https://github.com/Kavindu-Dodan/flagd/issues/863)) ([da30b7b](https://github.com/Kavindu-Dodan/flagd/commit/da30b7babffd8487c992fa41519787c8d78ebdba))
* add Health and Readiness probes ([#418](https://github.com/Kavindu-Dodan/flagd/issues/418)) ([7f2358c](https://github.com/Kavindu-Dodan/flagd/commit/7f2358ce207527c890f4a2f46ce4b9e8bf697095))
* add new configuration "sync-interval" which controls the HTTP polling interval ([#404](https://github.com/Kavindu-Dodan/flagd/issues/404)) ([ace62c7](https://github.com/Kavindu-Dodan/flagd/commit/ace62c7a6ab2b5b5d26642286deb6db406391d8f))
* Add version to startup message ([#430](https://github.com/Kavindu-Dodan/flagd/issues/430)) ([8daf613](https://github.com/Kavindu-Dodan/flagd/commit/8daf613e7e4f4492df0c06e2ef464f4337cadaca))
* expose Impression metric ([#556](https://github.com/Kavindu-Dodan/flagd/issues/556)) ([77e0a33](https://github.com/Kavindu-Dodan/flagd/commit/77e0a33be24dcd0b6e239e5ed709167167c14171))
* flagd image signing ([#338](https://github.com/Kavindu-Dodan/flagd/issues/338)) ([eca6a60](https://github.com/Kavindu-Dodan/flagd/commit/eca6a60967999a303ceef5465f1acc35c83afd6d))
* flagd OTEL collector ([#586](https://github.com/Kavindu-Dodan/flagd/issues/586)) ([494bec3](https://github.com/Kavindu-Dodan/flagd/commit/494bec33dcc1ddf0fa5cd0866f06265618408f5e))
* **flagd-proxy:** introduce zero-downtime ([#752](https://github.com/Kavindu-Dodan/flagd/issues/752)) ([ed5e6e5](https://github.com/Kavindu-Dodan/flagd/commit/ed5e6e5f3ee0a923c33dbf1a8bf20f80adec71bd))
* **flagd:** custom error handling for OTel errors ([#769](https://github.com/Kavindu-Dodan/flagd/issues/769)) ([bda1a92](https://github.com/Kavindu-Dodan/flagd/commit/bda1a92785c4348fe306a1d259b7bea91bd01c41))
* **flagD:** support zero downtime during upgrades ([#731](https://github.com/Kavindu-Dodan/flagd/issues/731)) ([7df8d39](https://github.com/Kavindu-Dodan/flagd/commit/7df8d3994b75991b5e49a65728ef5e4b24a85dde))
* grpc connection options to flagd configuration options ([#532](https://github.com/Kavindu-Dodan/flagd/issues/532)) ([aa74951](https://github.com/Kavindu-Dodan/flagd/commit/aa74951f43b662ff2df53e68d347fc10e6d23bb8))
* grpc selector as scope ([#761](https://github.com/Kavindu-Dodan/flagd/issues/761)) ([7246e6d](https://github.com/Kavindu-Dodan/flagd/commit/7246e6dce648c6445f90d71fc172bbab209d9928))
* grpc tls connectivity (grpcs) ([#477](https://github.com/Kavindu-Dodan/flagd/issues/477)) ([228f430](https://github.com/Kavindu-Dodan/flagd/commit/228f430e4945173755f52b8e712b23c28314517e))
* include falsy json fields ([#792](https://github.com/Kavindu-Dodan/flagd/issues/792)) ([37d91a0](https://github.com/Kavindu-Dodan/flagd/commit/37d91a09836f07e07b12acd13850ea5c7c9252cd))
* Introduce connect traces ([#624](https://github.com/Kavindu-Dodan/flagd/issues/624)) ([28bac6a](https://github.com/Kavindu-Dodan/flagd/commit/28bac6a54aed79cb8d84a147ffea296c36f5bd51))
* introduce flag merge behaviour ([#414](https://github.com/Kavindu-Dodan/flagd/issues/414)) ([524f65e](https://github.com/Kavindu-Dodan/flagd/commit/524f65ea7215466bb4ac24a8d0d5953dd1cfe9a0))
* Introduce flagd kube proxy ([#495](https://github.com/Kavindu-Dodan/flagd/issues/495)) ([440864c](https://github.com/Kavindu-Dodan/flagd/commit/440864ce87174618321c9d5146221490d8f07b24))
* introduce grpc sync for flagd ([#297](https://github.com/Kavindu-Dodan/flagd/issues/297)) ([33413f2](https://github.com/Kavindu-Dodan/flagd/commit/33413f25882a3f1cf4953da0f18e746bfb69faf4))
* Introduce kube-proxy-metrics ([#558](https://github.com/Kavindu-Dodan/flagd/issues/558)) ([ad0baeb](https://github.com/Kavindu-Dodan/flagd/commit/ad0baeb08fa67c94356d6a3f298283373bd5211b))
* introduce metrics for failed evaluations ([#584](https://github.com/Kavindu-Dodan/flagd/issues/584)) ([77664cd](https://github.com/Kavindu-Dodan/flagd/commit/77664cdf53a868f56ca040bdfe3f4930cd9a8fb4))
* introduce per-sync configurations ([#448](https://github.com/Kavindu-Dodan/flagd/issues/448)) ([1d80039](https://github.com/Kavindu-Dodan/flagd/commit/1d80039558b29fff117478e308fd794a1244f0e5))
* otel traces for flag evaluation ([#598](https://github.com/Kavindu-Dodan/flagd/issues/598)) ([1757035](https://github.com/Kavindu-Dodan/flagd/commit/175703548f88469f25d749e320ee48030c9f9074))
* proposal - rename `kube-flagd-proxy` to `flagd-proxy` ([#576](https://github.com/Kavindu-Dodan/flagd/issues/576)) ([223de99](https://github.com/Kavindu-Dodan/flagd/commit/223de99ee3efbcd601bf75ab1f6258eeac0c426e))
* refactor and improve K8s sync provider ([#443](https://github.com/Kavindu-Dodan/flagd/issues/443)) ([4c03bfc](https://github.com/Kavindu-Dodan/flagd/commit/4c03bfc812e7ceabcac0979290bd74d9efc9da15))
* refactor core module into multiple packages ([#530](https://github.com/Kavindu-Dodan/flagd/issues/530)) ([9d68d0b](https://github.com/Kavindu-Dodan/flagd/commit/9d68d0b45815facdf6079ffcd7864f720ccb8475))
* rename metrics and service ([#730](https://github.com/Kavindu-Dodan/flagd/issues/730)) ([09c0198](https://github.com/Kavindu-Dodan/flagd/commit/09c0198f76a200b1b6a1f48e9c94ec0547283ca2))
* Restructure for monorepo setup ([#486](https://github.com/Kavindu-Dodan/flagd/issues/486)) ([ed2993c](https://github.com/Kavindu-Dodan/flagd/commit/ed2993cd67b8a46db3beb6bb8a360e1aa20349da))
* support nested props in fractional evaluator ([#869](https://github.com/Kavindu-Dodan/flagd/issues/869)) ([50ff739](https://github.com/Kavindu-Dodan/flagd/commit/50ff739178fb732e38a220bb6a071260af1f2469))
* support OFO v1beta1 API ([#997](https://github.com/Kavindu-Dodan/flagd/issues/997)) ([bb6f5bf](https://github.com/Kavindu-Dodan/flagd/commit/bb6f5bf0fc382ade75d80a34d209beaa2edc459d))
* telemetry improvements ([#653](https://github.com/Kavindu-Dodan/flagd/issues/653)) ([ea02cba](https://github.com/Kavindu-Dodan/flagd/commit/ea02cba24bde982d55956fe54de1e8f27226bfc6))
* Use OTel to export metrics (metric name changes) ([#419](https://github.com/Kavindu-Dodan/flagd/issues/419)) ([eb3982a](https://github.com/Kavindu-Dodan/flagd/commit/eb3982a1cb72d664022b5cb126b533cf61497001))


### Bug Fixes

* add component tag ([#1050](https://github.com/Kavindu-Dodan/flagd/issues/1050)) ([361a301](https://github.com/Kavindu-Dodan/flagd/commit/361a301d51c034b4ed6dd92cd5061b889c8b2f97))
* add registry login ([#476](https://github.com/Kavindu-Dodan/flagd/issues/476)) ([99de755](https://github.com/Kavindu-Dodan/flagd/commit/99de755749df43d2b1028d47487b78b0ab626a9e))
* benchmark pipeline ([#538](https://github.com/Kavindu-Dodan/flagd/issues/538)) ([62cc0fc](https://github.com/Kavindu-Dodan/flagd/commit/62cc0fcfd6a63a6059352704117dbb78160eb689))
* **deps:** update dependency mkdocs to v1.5.3 ([#922](https://github.com/Kavindu-Dodan/flagd/issues/922)) ([dd7839e](https://github.com/Kavindu-Dodan/flagd/commit/dd7839e40562fb941b2375c7e64f15ecde4938d1))
* **deps:** update dependency mkdocs-material to v9.4.5 ([#852](https://github.com/Kavindu-Dodan/flagd/issues/852)) ([8a4de03](https://github.com/Kavindu-Dodan/flagd/commit/8a4de03375f7ccf5f60260901d79ba2e05d5b872))
* **deps:** update dependency mkdocs-material to v9.4.8 ([#969](https://github.com/Kavindu-Dodan/flagd/issues/969)) ([e6f92d2](https://github.com/Kavindu-Dodan/flagd/commit/e6f92d21089ba6e401bf58a3fef0c674fd783da4))
* **deps:** update dependency pillow to v10.1.0 ([#970](https://github.com/Kavindu-Dodan/flagd/issues/970)) ([47cd559](https://github.com/Kavindu-Dodan/flagd/commit/47cd559be4dd0b4e6baf554f0b9c1f892228a4f7))
* **deps:** update dependency pymdown-extensions to v10.3.1 ([#872](https://github.com/Kavindu-Dodan/flagd/issues/872)) ([3d2cd1f](https://github.com/Kavindu-Dodan/flagd/commit/3d2cd1fc0da4ee70abddbd1a6070b46daae050cc))
* **deps:** update dependency pymdown-extensions to v10.4 ([#1002](https://github.com/Kavindu-Dodan/flagd/issues/1002)) ([99a0fb5](https://github.com/Kavindu-Dodan/flagd/commit/99a0fb5ec7207ebec0f476450b3a6ec1b57b50ed))
* **deps:** update golang.org/x/exp digest to 9a3e603 ([#929](https://github.com/Kavindu-Dodan/flagd/issues/929)) ([f8db930](https://github.com/Kavindu-Dodan/flagd/commit/f8db930da22f52ffddd4533bdbaf8188f89250d0))
* **deps:** update kubernetes packages to v0.26.2 ([#450](https://github.com/Kavindu-Dodan/flagd/issues/450)) ([2885227](https://github.com/Kavindu-Dodan/flagd/commit/28852270f34ff81c072337b29aa17f4b6634e9cc))
* **deps:** update kubernetes packages to v0.26.3 ([#533](https://github.com/Kavindu-Dodan/flagd/issues/533)) ([6ddd5b2](https://github.com/Kavindu-Dodan/flagd/commit/6ddd5b29806f3101cf122bfc4196ba7d0ef4c025))
* **deps:** update kubernetes packages to v0.27.3 ([#708](https://github.com/Kavindu-Dodan/flagd/issues/708)) ([5bf3a69](https://github.com/Kavindu-Dodan/flagd/commit/5bf3a69aa4bf95ce77ad08491bcce420620525d3))
* **deps:** update kubernetes packages to v0.27.4 ([#756](https://github.com/Kavindu-Dodan/flagd/issues/756)) ([dcc10f3](https://github.com/Kavindu-Dodan/flagd/commit/dcc10f33f5fd9a8936241725ea811b90b4f136be))
* **deps:** update kubernetes packages to v0.28.0 ([#841](https://github.com/Kavindu-Dodan/flagd/issues/841)) ([cc195e1](https://github.com/Kavindu-Dodan/flagd/commit/cc195e1dde052d583656d5e5b49caec50f832365))
* **deps:** update kubernetes packages to v0.28.1 ([#860](https://github.com/Kavindu-Dodan/flagd/issues/860)) ([f3237c2](https://github.com/Kavindu-Dodan/flagd/commit/f3237c2d324fbb15fd5f7fe337a0601af3b537bb))
* **deps:** update kubernetes packages to v0.28.2 ([#911](https://github.com/Kavindu-Dodan/flagd/issues/911)) ([2eda6ab](https://github.com/Kavindu-Dodan/flagd/commit/2eda6ab5e528f12a9ce6b6818e08abb0d783b23d))
* **deps:** update kubernetes packages to v0.28.3 ([#974](https://github.com/Kavindu-Dodan/flagd/issues/974)) ([d7d205f](https://github.com/Kavindu-Dodan/flagd/commit/d7d205f457e46de3385610ee27db6dfd41323cd1))
* **deps:** update kubernetes packages to v0.28.4 ([#1016](https://github.com/Kavindu-Dodan/flagd/issues/1016)) ([ae470e3](https://github.com/Kavindu-Dodan/flagd/commit/ae470e37f3368c81484f0c54366ccf059ea2cea6))
* **deps:** update module buf.build/gen/go/open-feature/flagd/grpc/go to v1.3.0-20230720212818-3675556880a1.1 ([#747](https://github.com/Kavindu-Dodan/flagd/issues/747)) ([fb17bc6](https://github.com/Kavindu-Dodan/flagd/commit/fb17bc6a5c715f507b2838c150dc8a2f139a38fb))
* **deps:** update module github.com/bufbuild/connect-go to v1.10.0 ([#771](https://github.com/Kavindu-Dodan/flagd/issues/771)) ([c74103f](https://github.com/Kavindu-Dodan/flagd/commit/c74103faec068f14c87ad3ec227f5b802dbfac43))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.1 ([#365](https://github.com/Kavindu-Dodan/flagd/issues/365)) ([e25f452](https://github.com/Kavindu-Dodan/flagd/commit/e25f452906e034e339309270cc8db6dcd58e9973))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.2 ([#416](https://github.com/Kavindu-Dodan/flagd/issues/416)) ([feb7f04](https://github.com/Kavindu-Dodan/flagd/commit/feb7f047365263758a63d8dffea936f621a4966d))
* **deps:** update module github.com/bufbuild/connect-go to v1.6.0 ([#585](https://github.com/Kavindu-Dodan/flagd/issues/585)) ([8f2f467](https://github.com/Kavindu-Dodan/flagd/commit/8f2f467af52a3686196a821eec61954d89d3f71d))
* **deps:** update module github.com/bufbuild/connect-go to v1.7.0 ([#625](https://github.com/Kavindu-Dodan/flagd/issues/625)) ([1b24fc9](https://github.com/Kavindu-Dodan/flagd/commit/1b24fc923a405b337634009831ef0b9792953ce5))
* **deps:** update module github.com/bufbuild/connect-go to v1.8.0 ([#683](https://github.com/Kavindu-Dodan/flagd/issues/683)) ([13bb13d](https://github.com/Kavindu-Dodan/flagd/commit/13bb13daa11068481ba97f3432ae08de78392a91))
* **deps:** update module github.com/bufbuild/connect-go to v1.9.0 ([#722](https://github.com/Kavindu-Dodan/flagd/issues/722)) ([75223e2](https://github.com/Kavindu-Dodan/flagd/commit/75223e2fc01c4dcd0291b46a0d50b8815b31654c))
* **deps:** update module github.com/bufbuild/connect-opentelemetry-go to v0.3.0 ([#669](https://github.com/Kavindu-Dodan/flagd/issues/669)) ([e899435](https://github.com/Kavindu-Dodan/flagd/commit/e899435c29c32264ea2477436e69ce92c7775ee9))
* **deps:** update module github.com/bufbuild/connect-opentelemetry-go to v0.4.0 ([#739](https://github.com/Kavindu-Dodan/flagd/issues/739)) ([713e2a9](https://github.com/Kavindu-Dodan/flagd/commit/713e2a9834546963615046de1b6125e7fa6bf20d))
* **deps:** update module github.com/cucumber/godog to v0.13.0 ([#855](https://github.com/Kavindu-Dodan/flagd/issues/855)) ([5b42486](https://github.com/Kavindu-Dodan/flagd/commit/5b4248654f7199afc50663e73609eeb20a3d11ec))
* **deps:** update module github.com/diegoholiveira/jsonlogic/v3 to v3.3.0 ([#785](https://github.com/Kavindu-Dodan/flagd/issues/785)) ([ee9c54b](https://github.com/Kavindu-Dodan/flagd/commit/ee9c54b6b5cd51b947aae1ff6309ffae07ce89eb))
* **deps:** update module github.com/diegoholiveira/jsonlogic/v3 to v3.3.1 ([#971](https://github.com/Kavindu-Dodan/flagd/issues/971)) ([f1a40b8](https://github.com/Kavindu-Dodan/flagd/commit/f1a40b862e7be69b542ab65a645c29c56b8fe307))
* **deps:** update module github.com/diegoholiveira/jsonlogic/v3 to v3.3.2 ([#975](https://github.com/Kavindu-Dodan/flagd/issues/975)) ([b53c14a](https://github.com/Kavindu-Dodan/flagd/commit/b53c14afabd6b3f2403d9e2ea2bd26c8f1cfe8e4))
* **deps:** update module github.com/fsnotify/fsnotify to v1.7.0 ([#981](https://github.com/Kavindu-Dodan/flagd/issues/981)) ([727b9d2](https://github.com/Kavindu-Dodan/flagd/commit/727b9d2046ccd998d2ab721df10d88b84445494f))
* **deps:** update module github.com/open-feature/flagd/core to v0.4.5 ([#552](https://github.com/Kavindu-Dodan/flagd/issues/552)) ([41799f6](https://github.com/Kavindu-Dodan/flagd/commit/41799f624c261a84599cdd406cf28f4b33e49851))
* **deps:** update module github.com/open-feature/flagd/core to v0.5.1 ([#579](https://github.com/Kavindu-Dodan/flagd/issues/579)) ([58eed62](https://github.com/Kavindu-Dodan/flagd/commit/58eed62f5021e5c7a01a171067b725bf3ff83965))
* **deps:** update module github.com/open-feature/flagd/core to v0.5.2 ([#613](https://github.com/Kavindu-Dodan/flagd/issues/613)) ([218f435](https://github.com/Kavindu-Dodan/flagd/commit/218f435f0212fa24483b2af25e184e154e575eb1))
* **deps:** update module github.com/open-feature/flagd/core to v0.5.3 ([#634](https://github.com/Kavindu-Dodan/flagd/issues/634)) ([1bc7e99](https://github.com/Kavindu-Dodan/flagd/commit/1bc7e99473bc0c7bcacfb40030562e556d3895d6))
* **deps:** update module github.com/open-feature/flagd/core to v0.5.4 ([#693](https://github.com/Kavindu-Dodan/flagd/issues/693)) ([33705a6](https://github.com/Kavindu-Dodan/flagd/commit/33705a67300ec70760ba0baeb610f5a2e931205f))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.1 ([#745](https://github.com/Kavindu-Dodan/flagd/issues/745)) ([d290d8f](https://github.com/Kavindu-Dodan/flagd/commit/d290d8fda8aa84ed2db6454fdd26e60b028e3f7f))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.2 ([#779](https://github.com/Kavindu-Dodan/flagd/issues/779)) ([f34de59](https://github.com/Kavindu-Dodan/flagd/commit/f34de59fc8e636be043ce89758950d6ea3fe7376))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.3 ([#794](https://github.com/Kavindu-Dodan/flagd/issues/794)) ([9671964](https://github.com/Kavindu-Dodan/flagd/commit/96719649affeb1f8412e8b25f52d7292281d8230))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.4 ([#880](https://github.com/Kavindu-Dodan/flagd/issues/880)) ([ebb543d](https://github.com/Kavindu-Dodan/flagd/commit/ebb543d6eec18134e44ee7fe623fd2a336a1cf8d))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.5 ([#900](https://github.com/Kavindu-Dodan/flagd/issues/900)) ([c2ddcbf](https://github.com/Kavindu-Dodan/flagd/commit/c2ddcbfe49b8507fe463c11eb2b031bbc331792a))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.6 ([#916](https://github.com/Kavindu-Dodan/flagd/issues/916)) ([1f80e4d](https://github.com/Kavindu-Dodan/flagd/commit/1f80e4db9f8d1ba24884a71f2f8d552499ab5fe2))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.7 ([#966](https://github.com/Kavindu-Dodan/flagd/issues/966)) ([c038a3a](https://github.com/Kavindu-Dodan/flagd/commit/c038a3a3700eee82afa3e2cb2484614ec6ed566c))
* **deps:** update module github.com/open-feature/flagd/core to v0.6.8 ([#1006](https://github.com/Kavindu-Dodan/flagd/issues/1006)) ([c9b48bd](https://github.com/Kavindu-Dodan/flagd/commit/c9b48bd0b617f6d3c04c8924b1d6650ba17de81a))
* **deps:** update module github.com/open-feature/flagd/core to v0.7.0 ([#1014](https://github.com/Kavindu-Dodan/flagd/issues/1014)) ([deec49e](https://github.com/Kavindu-Dodan/flagd/commit/deec49e99ef52f62adbf278a8f58936acbb86b9d))
* **deps:** update module github.com/open-feature/flagd/core to v0.7.1 ([#1037](https://github.com/Kavindu-Dodan/flagd/issues/1037)) ([0ed9b68](https://github.com/Kavindu-Dodan/flagd/commit/0ed9b68341d026681c684a726b215ff910fe2a00))
* **deps:** update module github.com/open-feature/go-sdk to v1.8.0 ([#994](https://github.com/Kavindu-Dodan/flagd/issues/994)) ([266cf9f](https://github.com/Kavindu-Dodan/flagd/commit/266cf9f82ee8b4a4ba8ad1c0594388d2987a8c4b))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.10 ([#459](https://github.com/Kavindu-Dodan/flagd/issues/459)) ([cbdf9b0](https://github.com/Kavindu-Dodan/flagd/commit/cbdf9b07c30239d7d04ef770cf4461fb33422fe9))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.12 ([#635](https://github.com/Kavindu-Dodan/flagd/issues/635)) ([fe88061](https://github.com/Kavindu-Dodan/flagd/commit/fe88061ed6e0f1b6119af4c96a02495c4ff8072b))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.13 ([#697](https://github.com/Kavindu-Dodan/flagd/issues/697)) ([435448f](https://github.com/Kavindu-Dodan/flagd/commit/435448f449044eb5fff88c94e81883cc801c02c4))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.17 ([#759](https://github.com/Kavindu-Dodan/flagd/issues/759)) ([a2a2c3c](https://github.com/Kavindu-Dodan/flagd/commit/a2a2c3c7effd1708136eaac5df00ae02276d5005))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.18 ([#1011](https://github.com/Kavindu-Dodan/flagd/issues/1011)) ([90d4e4e](https://github.com/Kavindu-Dodan/flagd/commit/90d4e4e7d9db9e21fa38d96fdecb81ab78868732))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.9 ([#427](https://github.com/Kavindu-Dodan/flagd/issues/427)) ([42d2705](https://github.com/Kavindu-Dodan/flagd/commit/42d270558bf9badcff9a9b352fda35491c45aebe))
* **deps:** update module github.com/open-feature/go-sdk-contrib/tests/flagd to v1.2.2 ([#651](https://github.com/Kavindu-Dodan/flagd/issues/651)) ([9776973](https://github.com/Kavindu-Dodan/flagd/commit/9776973109a1bb45ab611ede6b2c4d2c01508455))
* **deps:** update module github.com/open-feature/go-sdk-contrib/tests/flagd to v1.2.3 ([#749](https://github.com/Kavindu-Dodan/flagd/issues/749)) ([cd63e48](https://github.com/Kavindu-Dodan/flagd/commit/cd63e489d681c0998a9c38072410653473ce40fc))
* **deps:** update module github.com/open-feature/go-sdk-contrib/tests/flagd to v1.3.1 ([#760](https://github.com/Kavindu-Dodan/flagd/issues/760)) ([30dda72](https://github.com/Kavindu-Dodan/flagd/commit/30dda72145c05de298140f880238ed37be73631a))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.29 ([#429](https://github.com/Kavindu-Dodan/flagd/issues/429)) ([b7fae81](https://github.com/Kavindu-Dodan/flagd/commit/b7fae81b89b3a1a0793a688c32569c4284633c6a))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.31 ([#527](https://github.com/Kavindu-Dodan/flagd/issues/527)) ([7928130](https://github.com/Kavindu-Dodan/flagd/commit/7928130b10906b10f4501630f16a71bdd8e4e365))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.32 [security] ([#606](https://github.com/Kavindu-Dodan/flagd/issues/606)) ([6f721af](https://github.com/Kavindu-Dodan/flagd/commit/6f721af379fcb0f1a74410637a313477148ef863))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.34 ([#604](https://github.com/Kavindu-Dodan/flagd/issues/604)) ([3e6a84b](https://github.com/Kavindu-Dodan/flagd/commit/3e6a84b455a330f541784b346d3b6199f8b423f7))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.35 ([#783](https://github.com/Kavindu-Dodan/flagd/issues/783)) ([9ff0b5b](https://github.com/Kavindu-Dodan/flagd/commit/9ff0b5b1bd3bb95581eab83f944aa60e179b207a))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.36 ([#799](https://github.com/Kavindu-Dodan/flagd/issues/799)) ([fa4da4b](https://github.com/Kavindu-Dodan/flagd/commit/fa4da4b0115e9fb40ab038b996e1e32b9f6a47ab))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.5.0 ([#1039](https://github.com/Kavindu-Dodan/flagd/issues/1039)) ([eb128d9](https://github.com/Kavindu-Dodan/flagd/commit/eb128d97ecb8b6916f9c5a32c21c698207f82be5))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.5.1 ([#1046](https://github.com/Kavindu-Dodan/flagd/issues/1046)) ([0321935](https://github.com/Kavindu-Dodan/flagd/commit/0321935992ef50c2fc62bd659d8858c25c88efa7))
* **deps:** update module github.com/prometheus/client_golang to v1.15.0 ([#608](https://github.com/Kavindu-Dodan/flagd/issues/608)) ([0597a8f](https://github.com/Kavindu-Dodan/flagd/commit/0597a8f23d0914b26f06b1335b49fe7c18ecb4f9))
* **deps:** update module github.com/prometheus/client_golang to v1.15.1 ([#636](https://github.com/Kavindu-Dodan/flagd/issues/636)) ([b22279d](https://github.com/Kavindu-Dodan/flagd/commit/b22279df469dc78f9d3e5bc4a59ab6baf539a8ae))
* **deps:** update module github.com/prometheus/client_golang to v1.16.0 ([#709](https://github.com/Kavindu-Dodan/flagd/issues/709)) ([b8bedd2](https://github.com/Kavindu-Dodan/flagd/commit/b8bedd2b895026eace8204ae4ffcff771f7e8e97))
* **deps:** update module github.com/prometheus/client_golang to v1.17.0 ([#939](https://github.com/Kavindu-Dodan/flagd/issues/939)) ([9065cba](https://github.com/Kavindu-Dodan/flagd/commit/9065cba599ac07225b613c5acd9403ab24462078))
* **deps:** update module github.com/rs/cors to v1.10.0 ([#893](https://github.com/Kavindu-Dodan/flagd/issues/893)) ([fe61fbe](https://github.com/Kavindu-Dodan/flagd/commit/fe61fbe47a4e58562cbcb1c5201281fae1adafaf))
* **deps:** update module github.com/rs/cors to v1.10.1 ([#946](https://github.com/Kavindu-Dodan/flagd/issues/946)) ([1c39862](https://github.com/Kavindu-Dodan/flagd/commit/1c39862297746e66189ece87892b6e4694294fb6))
* **deps:** update module github.com/rs/cors to v1.9.0 ([#609](https://github.com/Kavindu-Dodan/flagd/issues/609)) ([97066c1](https://github.com/Kavindu-Dodan/flagd/commit/97066c107d14777eaad8a05b3bb051639af3179c))
* **deps:** update module github.com/rs/xid to v1.5.0 ([#614](https://github.com/Kavindu-Dodan/flagd/issues/614)) ([e3dfbc6](https://github.com/Kavindu-Dodan/flagd/commit/e3dfbc6753bddc9e2f5c4a2633a7010123cf2f97))
* **deps:** update module github.com/spf13/cobra to v1.7.0 ([#587](https://github.com/Kavindu-Dodan/flagd/issues/587)) ([12b3477](https://github.com/Kavindu-Dodan/flagd/commit/12b34773a68f6ae7e7e605aebc9f7075eb819994))
* **deps:** update module github.com/spf13/cobra to v1.8.0 ([#993](https://github.com/Kavindu-Dodan/flagd/issues/993)) ([05c7870](https://github.com/Kavindu-Dodan/flagd/commit/05c7870cc7662117f85e9c6528508327ae320b83))
* **deps:** update module github.com/spf13/viper to v1.16.0 ([#679](https://github.com/Kavindu-Dodan/flagd/issues/679)) ([798a975](https://github.com/Kavindu-Dodan/flagd/commit/798a975bb1a47420e814b6dd439f1cece1a263e5))
* **deps:** update module github.com/spf13/viper to v1.17.0 ([#956](https://github.com/Kavindu-Dodan/flagd/issues/956)) ([31d015d](https://github.com/Kavindu-Dodan/flagd/commit/31d015d329ae9c1da3ec13878078371bcbf43fbf))
* **deps:** update module github.com/stretchr/testify to v1.8.2 ([#440](https://github.com/Kavindu-Dodan/flagd/issues/440)) ([ab3e674](https://github.com/Kavindu-Dodan/flagd/commit/ab3e6748abc7843c022afeaf7cb11193cdcf59c5))
* **deps:** update module github.com/stretchr/testify to v1.8.3 ([#662](https://github.com/Kavindu-Dodan/flagd/issues/662)) ([2e06d58](https://github.com/Kavindu-Dodan/flagd/commit/2e06d582ee9c8abfd57f8945d91261eab6cf9854))
* **deps:** update module github.com/stretchr/testify to v1.8.4 ([#678](https://github.com/Kavindu-Dodan/flagd/issues/678)) ([ca8c9d6](https://github.com/Kavindu-Dodan/flagd/commit/ca8c9d66a0c6b21129c4c36a3c10dcf3be869ee7))
* **deps:** update module go.uber.org/zap to v1.25.0 ([#786](https://github.com/Kavindu-Dodan/flagd/issues/786)) ([40d0aa6](https://github.com/Kavindu-Dodan/flagd/commit/40d0aa66cf422db6811206d777b55396a96f330f))
* **deps:** update module go.uber.org/zap to v1.26.0 ([#917](https://github.com/Kavindu-Dodan/flagd/issues/917)) ([e57e206](https://github.com/Kavindu-Dodan/flagd/commit/e57e206c937d5b11b81d46ee57b3e92cc454dd88))
* **deps:** update module golang.org/x/crypto to v0.10.0 ([#647](https://github.com/Kavindu-Dodan/flagd/issues/647)) ([7f1d7e6](https://github.com/Kavindu-Dodan/flagd/commit/7f1d7e66669b88b2c56b32f9cdd9be354ebcfc8e))
* **deps:** update module golang.org/x/crypto to v0.12.0 ([#797](https://github.com/Kavindu-Dodan/flagd/issues/797)) ([edae3fd](https://github.com/Kavindu-Dodan/flagd/commit/edae3fd466c0be62a0256c268e85cb337c9536f2))
* **deps:** update module golang.org/x/crypto to v0.13.0 ([#888](https://github.com/Kavindu-Dodan/flagd/issues/888)) ([1a9037a](https://github.com/Kavindu-Dodan/flagd/commit/1a9037a5b058e44fa844392d0110696b032eff6e))
* **deps:** update module golang.org/x/crypto to v0.16.0 ([#1033](https://github.com/Kavindu-Dodan/flagd/issues/1033)) ([b79aaf2](https://github.com/Kavindu-Dodan/flagd/commit/b79aaf2c2ca5fe9c43c1460d6cab349d5d68b7e1))
* **deps:** update module golang.org/x/crypto to v0.7.0 ([#472](https://github.com/Kavindu-Dodan/flagd/issues/472)) ([f53f6c8](https://github.com/Kavindu-Dodan/flagd/commit/f53f6c885ee90813161b99be5a273b485e064de8))
* **deps:** update module golang.org/x/crypto to v0.8.0 ([#595](https://github.com/Kavindu-Dodan/flagd/issues/595)) ([36016d7](https://github.com/Kavindu-Dodan/flagd/commit/36016d7940fa772c01dd61b071b2c9ec753cfb75))
* **deps:** update module golang.org/x/mod to v0.10.0 ([#682](https://github.com/Kavindu-Dodan/flagd/issues/682)) ([16199ce](https://github.com/Kavindu-Dodan/flagd/commit/16199ceac9ebbae68dafbd6c21239f64f8c32511))
* **deps:** update module golang.org/x/mod to v0.11.0 ([#705](https://github.com/Kavindu-Dodan/flagd/issues/705)) ([42813be](https://github.com/Kavindu-Dodan/flagd/commit/42813bef092ba7fffed0dd94166bfd01ea8a7582))
* **deps:** update module golang.org/x/mod to v0.12.0 ([#729](https://github.com/Kavindu-Dodan/flagd/issues/729)) ([7b109c7](https://github.com/Kavindu-Dodan/flagd/commit/7b109c705aceb652ac2675bd0ffe82420983798b))
* **deps:** update module golang.org/x/mod to v0.13.0 ([#952](https://github.com/Kavindu-Dodan/flagd/issues/952)) ([be61450](https://github.com/Kavindu-Dodan/flagd/commit/be61450ec3fab3c294c9813df193c98e374900aa))
* **deps:** update module golang.org/x/mod to v0.14.0 ([#991](https://github.com/Kavindu-Dodan/flagd/issues/991)) ([87bc12d](https://github.com/Kavindu-Dodan/flagd/commit/87bc12dd968e9fcf524eb0b69462222f29c6ba34))
* **deps:** update module golang.org/x/net to v0.10.0 ([#644](https://github.com/Kavindu-Dodan/flagd/issues/644)) ([ccd9d35](https://github.com/Kavindu-Dodan/flagd/commit/ccd9d351df153039a124064f30e5829610773f27))
* **deps:** update module golang.org/x/net to v0.11.0 ([#706](https://github.com/Kavindu-Dodan/flagd/issues/706)) ([27d893f](https://github.com/Kavindu-Dodan/flagd/commit/27d893fe78417f7b8418003edc401ab5a6c21fb9))
* **deps:** update module golang.org/x/net to v0.12.0 ([#734](https://github.com/Kavindu-Dodan/flagd/issues/734)) ([777b28b](https://github.com/Kavindu-Dodan/flagd/commit/777b28b1d512245b0046d11197f6dfa341b317d2))
* **deps:** update module golang.org/x/net to v0.13.0 ([#784](https://github.com/Kavindu-Dodan/flagd/issues/784)) ([f57d023](https://github.com/Kavindu-Dodan/flagd/commit/f57d023174d9cc74b7d8260055f82b84a2bdcc52))
* **deps:** update module golang.org/x/net to v0.14.0 ([#798](https://github.com/Kavindu-Dodan/flagd/issues/798)) ([92c2f26](https://github.com/Kavindu-Dodan/flagd/commit/92c2f2676163688130737b34a115374cb5631247))
* **deps:** update module golang.org/x/net to v0.15.0 ([#889](https://github.com/Kavindu-Dodan/flagd/issues/889)) ([233d976](https://github.com/Kavindu-Dodan/flagd/commit/233d97694826d0e018be19a78259188802aba37f))
* **deps:** update module golang.org/x/net to v0.17.0 [security] ([#963](https://github.com/Kavindu-Dodan/flagd/issues/963)) ([7f54bd1](https://github.com/Kavindu-Dodan/flagd/commit/7f54bd1fb3fdbb7dff3a7b097f804ce843bb6e3a))
* **deps:** update module golang.org/x/net to v0.18.0 ([#1000](https://github.com/Kavindu-Dodan/flagd/issues/1000)) ([e9347cc](https://github.com/Kavindu-Dodan/flagd/commit/e9347cc88a4174c984432571710588ac57665cb7))
* **deps:** update module golang.org/x/net to v0.19.0 ([#1034](https://github.com/Kavindu-Dodan/flagd/issues/1034)) ([c6426b2](https://github.com/Kavindu-Dodan/flagd/commit/c6426b20fa241066fc844a2a000812015790cf12))
* **deps:** update module golang.org/x/net to v0.6.0 ([#396](https://github.com/Kavindu-Dodan/flagd/issues/396)) ([beb7564](https://github.com/Kavindu-Dodan/flagd/commit/beb756470b1e1d5ef0670b8322b6ed9cb44efa24))
* **deps:** update module golang.org/x/net to v0.7.0 ([#410](https://github.com/Kavindu-Dodan/flagd/issues/410)) ([c6133b6](https://github.com/Kavindu-Dodan/flagd/commit/c6133b6af61f3d73ae73d318a1a9f44db2540540))
* **deps:** update module golang.org/x/net to v0.8.0 ([#468](https://github.com/Kavindu-Dodan/flagd/issues/468)) ([10d5f2c](https://github.com/Kavindu-Dodan/flagd/commit/10d5f2c55081a25daa1f0e0fa81f96f0fffbbc7b))
* **deps:** update module golang.org/x/sync to v0.2.0 ([#638](https://github.com/Kavindu-Dodan/flagd/issues/638)) ([7f4a7db](https://github.com/Kavindu-Dodan/flagd/commit/7f4a7db8139294a21b3415710c143f182d93264a))
* **deps:** update module golang.org/x/sync to v0.3.0 ([#707](https://github.com/Kavindu-Dodan/flagd/issues/707)) ([7852efb](https://github.com/Kavindu-Dodan/flagd/commit/7852efb84e9f071b2b482b1968d799888b6882dc))
* **deps:** update module golang.org/x/sync to v0.4.0 ([#949](https://github.com/Kavindu-Dodan/flagd/issues/949)) ([faa24a6](https://github.com/Kavindu-Dodan/flagd/commit/faa24a6c6f34330364e1ba0c3a847943f4e55150))
* **deps:** update module golang.org/x/sync to v0.5.0 ([#992](https://github.com/Kavindu-Dodan/flagd/issues/992)) ([bd24536](https://github.com/Kavindu-Dodan/flagd/commit/bd24536722f3c3c99b5946de0b62dba603992ba7))
* **deps:** update module google.golang.org/grpc to v1.53.0 ([#388](https://github.com/Kavindu-Dodan/flagd/issues/388)) ([174cd7c](https://github.com/Kavindu-Dodan/flagd/commit/174cd7c70fa5ae2573db2c5972b75786633e2f41))
* **deps:** update module google.golang.org/grpc to v1.54.0 ([#548](https://github.com/Kavindu-Dodan/flagd/issues/548)) ([99ba5ec](https://github.com/Kavindu-Dodan/flagd/commit/99ba5ece76d98124c108bc6280bee03a5c0cd25d))
* **deps:** update module google.golang.org/grpc to v1.55.0 ([#640](https://github.com/Kavindu-Dodan/flagd/issues/640)) ([c0d7328](https://github.com/Kavindu-Dodan/flagd/commit/c0d732866262240e340fe10f8ac0f6ff2a5c4f8c))
* **deps:** update module google.golang.org/grpc to v1.56.1 ([#710](https://github.com/Kavindu-Dodan/flagd/issues/710)) ([8f16573](https://github.com/Kavindu-Dodan/flagd/commit/8f165739aee8f28800e200b357203e88a3fd5938))
* **deps:** update module google.golang.org/grpc to v1.56.2 ([#738](https://github.com/Kavindu-Dodan/flagd/issues/738)) ([521cc30](https://github.com/Kavindu-Dodan/flagd/commit/521cc30cde1971be000ec10d93f6d70b9b2260ee))
* **deps:** update module google.golang.org/grpc to v1.57.0 ([#773](https://github.com/Kavindu-Dodan/flagd/issues/773)) ([be8bf04](https://github.com/Kavindu-Dodan/flagd/commit/be8bf045093d89099eead2cccb86a5a7275e25d5))
* **deps:** update module google.golang.org/grpc to v1.58.0 ([#896](https://github.com/Kavindu-Dodan/flagd/issues/896)) ([853b76d](https://github.com/Kavindu-Dodan/flagd/commit/853b76dfa3babfebd8bdbcd3e0913380f077b8ab))
* **deps:** update module google.golang.org/grpc to v1.58.1 ([#915](https://github.com/Kavindu-Dodan/flagd/issues/915)) ([06d95de](https://github.com/Kavindu-Dodan/flagd/commit/06d95ded9b69c9c598d08f8a6ef73ec598a817af))
* **deps:** update module google.golang.org/grpc to v1.58.2 ([#928](https://github.com/Kavindu-Dodan/flagd/issues/928)) ([90f1878](https://github.com/Kavindu-Dodan/flagd/commit/90f1878ae482ea4a684615f733392a38301de68d))
* **deps:** update module google.golang.org/grpc to v1.58.3 ([#960](https://github.com/Kavindu-Dodan/flagd/issues/960)) ([fee1558](https://github.com/Kavindu-Dodan/flagd/commit/fee1558da4f5418499fb09fe356d16f008423eb7))
* **deps:** update module google.golang.org/grpc to v1.59.0 ([#972](https://github.com/Kavindu-Dodan/flagd/issues/972)) ([7d0f1f2](https://github.com/Kavindu-Dodan/flagd/commit/7d0f1f225ebfe83681c83cf38bf5c253a7b0ebef))
* **deps:** update module google.golang.org/protobuf to v1.29.0 ([#478](https://github.com/Kavindu-Dodan/flagd/issues/478)) ([f9adc8e](https://github.com/Kavindu-Dodan/flagd/commit/f9adc8e3746256bcec045c06c78034c45722d60c))
* **deps:** update module google.golang.org/protobuf to v1.29.1 [security] ([#504](https://github.com/Kavindu-Dodan/flagd/issues/504)) ([59db0bb](https://github.com/Kavindu-Dodan/flagd/commit/59db0bba43a9c002378fdced2fcf4729d619e28b))
* **deps:** update module google.golang.org/protobuf to v1.30.0 ([#499](https://github.com/Kavindu-Dodan/flagd/issues/499)) ([f650338](https://github.com/Kavindu-Dodan/flagd/commit/f650338e01e721a9d24e2ed6f6fe585dbb6beb42))
* **deps:** update module google.golang.org/protobuf to v1.31.0 ([#720](https://github.com/Kavindu-Dodan/flagd/issues/720)) ([247239e](https://github.com/Kavindu-Dodan/flagd/commit/247239e76b9de1a619aad9e957ed8b44ae534b77))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.3 ([#372](https://github.com/Kavindu-Dodan/flagd/issues/372)) ([330ac91](https://github.com/Kavindu-Dodan/flagd/commit/330ac91e375124826b2a7a1a22d0daa18368ab99))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.4 ([#374](https://github.com/Kavindu-Dodan/flagd/issues/374)) ([d90e561](https://github.com/Kavindu-Dodan/flagd/commit/d90e561bfc5b798d13d4ba8f30f523b1053f3748))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.5 ([#454](https://github.com/Kavindu-Dodan/flagd/issues/454)) ([f907f11](https://github.com/Kavindu-Dodan/flagd/commit/f907f114f23fa2efa2637e254e829e4d53a90b51))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.6 ([#572](https://github.com/Kavindu-Dodan/flagd/issues/572)) ([bed9458](https://github.com/Kavindu-Dodan/flagd/commit/bed94584a30bb6752284ece152cb114a102cbe8a))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.15.0 ([#665](https://github.com/Kavindu-Dodan/flagd/issues/665)) ([9490ed6](https://github.com/Kavindu-Dodan/flagd/commit/9490ed62e2fc589af8ae7ee26bfd559797a1f83c))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.15.1 ([#795](https://github.com/Kavindu-Dodan/flagd/issues/795)) ([13d62fd](https://github.com/Kavindu-Dodan/flagd/commit/13d62fd0fc4749f19dba0a18e1fda46a723380c5))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.0 ([#856](https://github.com/Kavindu-Dodan/flagd/issues/856)) ([88d832a](https://github.com/Kavindu-Dodan/flagd/commit/88d832a9d49a4bc1d6156849a59227ecab07f96e))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.1 ([#882](https://github.com/Kavindu-Dodan/flagd/issues/882)) ([ca3d85a](https://github.com/Kavindu-Dodan/flagd/commit/ca3d85a51c0ed1c1def54d7304d4b9fe69622662))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.2 ([#907](https://github.com/Kavindu-Dodan/flagd/issues/907)) ([9976851](https://github.com/Kavindu-Dodan/flagd/commit/9976851d792ff3eb5fde18f19e397738eb7cacaf))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.16.3 ([#976](https://github.com/Kavindu-Dodan/flagd/issues/976)) ([b33c9c9](https://github.com/Kavindu-Dodan/flagd/commit/b33c9c97cf21eb317e4b0b1f452571d85de145b1))
* **deps:** update opentelemetry-go monorepo ([#1001](https://github.com/Kavindu-Dodan/flagd/issues/1001)) ([9798aeb](https://github.com/Kavindu-Dodan/flagd/commit/9798aeb248764400128048b65ba27baba71b07e6))
* **deps:** update opentelemetry-go monorepo ([#1019](https://github.com/Kavindu-Dodan/flagd/issues/1019)) ([23ae555](https://github.com/Kavindu-Dodan/flagd/commit/23ae555ad73128dff46e911fccfef76306f5c550))
* **deps:** update opentelemetry-go monorepo ([#648](https://github.com/Kavindu-Dodan/flagd/issues/648)) ([c12dad8](https://github.com/Kavindu-Dodan/flagd/commit/c12dad89a8e761154f57739ded594b2783a14f8a))
* **deps:** update opentelemetry-go monorepo ([#868](https://github.com/Kavindu-Dodan/flagd/issues/868)) ([d48317f](https://github.com/Kavindu-Dodan/flagd/commit/d48317f61d7db7ba0398dc9ab7cdd174a0b87555))
* **deps:** update opentelemetry-go monorepo ([#906](https://github.com/Kavindu-Dodan/flagd/issues/906)) ([5a41226](https://github.com/Kavindu-Dodan/flagd/commit/5a4122658039aafcf080fcc6655c2a679622ed69))
* **deps:** update opentelemetry-go monorepo ([#943](https://github.com/Kavindu-Dodan/flagd/issues/943)) ([e7cee41](https://github.com/Kavindu-Dodan/flagd/commit/e7cee41630e5684999b3689dbf1ab66234c65f6e))
* erroneous warning about prop overwrite ([#924](https://github.com/Kavindu-Dodan/flagd/issues/924)) ([673b76a](https://github.com/Kavindu-Dodan/flagd/commit/673b76aeff5c27e8e031e59da5f0ed1871d3f749))
* error handling of Serve/ServeTLS funcs ([#397](https://github.com/Kavindu-Dodan/flagd/issues/397)) ([8923bf2](https://github.com/Kavindu-Dodan/flagd/commit/8923bf2d407e18b65c188aef9bf7370fc74c3be2))
* eventing configuration setup ([#605](https://github.com/Kavindu-Dodan/flagd/issues/605)) ([edfbe51](https://github.com/Kavindu-Dodan/flagd/commit/edfbe5191651f25da991b507a3feedcbbe3c66f1))
* fall back to default port if env var cannot be parsed ([#591](https://github.com/Kavindu-Dodan/flagd/issues/591)) ([1fda104](https://github.com/Kavindu-Dodan/flagd/commit/1fda10473dba36149e13fa0cb8bb686d6861e568))
* fix broken image signing ([#461](https://github.com/Kavindu-Dodan/flagd/issues/461)) ([05bb51c](https://github.com/Kavindu-Dodan/flagd/commit/05bb51c7ab30f6e976b87f54ca889e978f834211))
* fix connect error code handling for disabled flags ([#670](https://github.com/Kavindu-Dodan/flagd/issues/670)) ([86a8012](https://github.com/Kavindu-Dodan/flagd/commit/86a8012efcfeb3e967657f6143c143b457d64ca2))
* fix race in http sync test ([#401](https://github.com/Kavindu-Dodan/flagd/issues/401)) ([1d0c8e1](https://github.com/Kavindu-Dodan/flagd/commit/1d0c8e168b73f7fbd4b27ece733041bbe08261c0))
* fixing image delimeter  ([#463](https://github.com/Kavindu-Dodan/flagd/issues/463)) ([b4ee495](https://github.com/Kavindu-Dodan/flagd/commit/b4ee495dc8e00b032518ea42d272a36b3b662e95))
* fixing silent lint failures ([#550](https://github.com/Kavindu-Dodan/flagd/issues/550)) ([30c8022](https://github.com/Kavindu-Dodan/flagd/commit/30c8022e891d1d278c096dd2438137ced7552678))
* flagd-proxy locking bug fix ([#592](https://github.com/Kavindu-Dodan/flagd/issues/592)) ([b166122](https://github.com/Kavindu-Dodan/flagd/commit/b1661225c912ee11ba4749f7ef157a0335e8781f))
* init workspace before running scheduled benchmark tests ([#564](https://github.com/Kavindu-Dodan/flagd/issues/564)) ([cb3a3e8](https://github.com/Kavindu-Dodan/flagd/commit/cb3a3e84bb370db06be59b05e530d45353a0dad3))
* introduced RWMutex to flag state to prevent concurrent r/w of map ([#370](https://github.com/Kavindu-Dodan/flagd/issues/370)) ([93e356b](https://github.com/Kavindu-Dodan/flagd/commit/93e356b4ab0b65c71659bd52d73f618edffc96f5))
* metric descriptions match the otel spec ([#789](https://github.com/Kavindu-Dodan/flagd/issues/789)) ([34befcd](https://github.com/Kavindu-Dodan/flagd/commit/34befcdfedc5f0479cb0ae77fe148849c341d33e))
* nil pointer fix + export constructors ([#555](https://github.com/Kavindu-Dodan/flagd/issues/555)) ([78adb81](https://github.com/Kavindu-Dodan/flagd/commit/78adb81f4eb7a5b7fdb7075fa0bf8afa6d03dc72))
* remove disabled flags from bulk evaluation ([#672](https://github.com/Kavindu-Dodan/flagd/issues/672)) ([d2ce988](https://github.com/Kavindu-Dodan/flagd/commit/d2ce98838edf63b88ee9fb5ae6f8d534e1112e7e))
* remove non-error error log from parseFractionalEvaluationData ([#446](https://github.com/Kavindu-Dodan/flagd/issues/446)) ([34aca79](https://github.com/Kavindu-Dodan/flagd/commit/34aca79e6ec9876a6cced0fe49e1ceea34d83696))
* sbom artefact name ([#380](https://github.com/Kavindu-Dodan/flagd/issues/380)) ([3daef26](https://github.com/Kavindu-Dodan/flagd/commit/3daef263c43ed63776d604d27f7ae6b993fff143)), closes [#379](https://github.com/Kavindu-Dodan/flagd/issues/379)
* security issues ([#464](https://github.com/Kavindu-Dodan/flagd/issues/464)) ([7f1e759](https://github.com/Kavindu-Dodan/flagd/commit/7f1e759a87a9af63e9384005c959a3f500cc474c))
* set readiness once only ([#465](https://github.com/Kavindu-Dodan/flagd/issues/465)) ([41a888d](https://github.com/Kavindu-Dodan/flagd/commit/41a888d6b60c030b913280c2a1eeff8b25e8aada))
* set ResolveObject reason ([#375](https://github.com/Kavindu-Dodan/flagd/issues/375)) ([dcf199d](https://github.com/Kavindu-Dodan/flagd/commit/dcf199dab9d11b86454028869a54d77a474fc4a6))
* Set the formula path to match the restructured homebrew repo ([#881](https://github.com/Kavindu-Dodan/flagd/issues/881)) ([5048a97](https://github.com/Kavindu-Dodan/flagd/commit/5048a9792650f45b62305e82f475424f75f6d98d))
* use 32bit murmur calculation (64 is not stable) ([#913](https://github.com/Kavindu-Dodan/flagd/issues/913)) ([db8dca4](https://github.com/Kavindu-Dodan/flagd/commit/db8dca421cb0dba2968d47e5cc162d81401298db))


### Miscellaneous Chores

* unify sources configuration handling ([#560](https://github.com/Kavindu-Dodan/flagd/issues/560)) ([7f4888a](https://github.com/Kavindu-Dodan/flagd/commit/7f4888a1676e49acecf328685623566ea057ffcf))

## [0.4.2](https://github.com/open-feature/flagd/compare/v0.4.1...v0.4.2) (2023-03-09)


### 🧹 Chore

* Add targeted Flag to example config ([#467](https://github.com/open-feature/flagd/issues/467)) ([6a039ce](https://github.com/open-feature/flagd/commit/6a039cef875caae61ea6c65799f3b6dc3863d131))
* **deps:** pin dependencies ([#473](https://github.com/open-feature/flagd/issues/473)) ([679e860](https://github.com/open-feature/flagd/commit/679e8600f57ab1e03c493c4a4046bd9d7368efac))
* **deps:** update google-github-actions/release-please-action digest to e0b9d18 ([#474](https://github.com/open-feature/flagd/issues/474)) ([5b85b2a](https://github.com/open-feature/flagd/commit/5b85b2a611d9199e39735f101ed7e560257ce2e4))
* refactoring and improve coverage for K8s Sync ([#466](https://github.com/open-feature/flagd/issues/466)) ([6dc441e](https://github.com/open-feature/flagd/commit/6dc441e2f2418c1fd3a5a58dbb99f848ccbd8735))


### 🐛 Bug Fixes

* add registry login ([#476](https://github.com/open-feature/flagd/issues/476)) ([99de755](https://github.com/open-feature/flagd/commit/99de755749df43d2b1028d47487b78b0ab626a9e))
* **deps:** update module golang.org/x/crypto to v0.7.0 ([#472](https://github.com/open-feature/flagd/issues/472)) ([f53f6c8](https://github.com/open-feature/flagd/commit/f53f6c885ee90813161b99be5a273b485e064de8))
* **deps:** update module google.golang.org/protobuf to v1.29.0 ([#478](https://github.com/open-feature/flagd/issues/478)) ([f9adc8e](https://github.com/open-feature/flagd/commit/f9adc8e3746256bcec045c06c78034c45722d60c))


### ✨ New Features

* grpc tls connectivity (grpcs) ([#477](https://github.com/open-feature/flagd/issues/477)) ([228f430](https://github.com/open-feature/flagd/commit/228f430e4945173755f52b8e712b23c28314517e))
* introduce per-sync configurations ([#448](https://github.com/open-feature/flagd/issues/448)) ([1d80039](https://github.com/open-feature/flagd/commit/1d80039558b29fff117478e308fd794a1244f0e5))

## [0.4.1](https://github.com/open-feature/flagd/compare/v0.4.0...v0.4.1) (2023-03-07)


### 🔄 Refactoring

* remove unused struct field ([#458](https://github.com/open-feature/flagd/issues/458)) ([a04c0b8](https://github.com/open-feature/flagd/commit/a04c0b837dbe9e28d1e01e43ea9e378a6c0f316a))


### 🧹 Chore

* **deps:** update sigstore/cosign-installer digest to bd2d118 ([#471](https://github.com/open-feature/flagd/issues/471)) ([ee90f48](https://github.com/open-feature/flagd/commit/ee90f48317ec600f09534306503dc752254a1d09))


### 🐛 Bug Fixes

* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.10 ([#459](https://github.com/open-feature/flagd/issues/459)) ([cbdf9b0](https://github.com/open-feature/flagd/commit/cbdf9b07c30239d7d04ef770cf4461fb33422fe9))
* **deps:** update module golang.org/x/net to v0.8.0 ([#468](https://github.com/open-feature/flagd/issues/468)) ([10d5f2c](https://github.com/open-feature/flagd/commit/10d5f2c55081a25daa1f0e0fa81f96f0fffbbc7b))
* fix broken image signing ([#461](https://github.com/open-feature/flagd/issues/461)) ([05bb51c](https://github.com/open-feature/flagd/commit/05bb51c7ab30f6e976b87f54ca889e978f834211))
* fixing image delimeter  ([#463](https://github.com/open-feature/flagd/issues/463)) ([b4ee495](https://github.com/open-feature/flagd/commit/b4ee495dc8e00b032518ea42d272a36b3b662e95))
* security issues ([#464](https://github.com/open-feature/flagd/issues/464)) ([7f1e759](https://github.com/open-feature/flagd/commit/7f1e759a87a9af63e9384005c959a3f500cc474c))
* set readiness once only ([#465](https://github.com/open-feature/flagd/issues/465)) ([41a888d](https://github.com/open-feature/flagd/commit/41a888d6b60c030b913280c2a1eeff8b25e8aada))

## [0.4.0](https://github.com/open-feature/flagd/compare/v0.3.7...v0.4.0) (2023-03-02)


### ⚠ BREAKING CHANGES

* Use OTel to export metrics (metric name changes) ([#419](https://github.com/open-feature/flagd/issues/419))

### 🧹 Chore

* add additional sections to the release notes ([#449](https://github.com/open-feature/flagd/issues/449)) ([798f71a](https://github.com/open-feature/flagd/commit/798f71a92d2e2f450a53cda93b44217cbb2ad7fd))
* attach image sbom to release artefacts ([#407](https://github.com/open-feature/flagd/issues/407)) ([fb4ee50](https://github.com/open-feature/flagd/commit/fb4ee502217e2262849df09258f3a0ffa7edec13))
* **deps:** update actions/configure-pages digest to fc89b04 ([#417](https://github.com/open-feature/flagd/issues/417)) ([04014e7](https://github.com/open-feature/flagd/commit/04014e7cb37e43f5ed3726dfd31da96202abc043))
* **deps:** update amannn/action-semantic-pull-request digest to b6bca70 ([#441](https://github.com/open-feature/flagd/issues/441)) ([ce0ebe1](https://github.com/open-feature/flagd/commit/ce0ebe13dd992688a3a0464ff401f2c40651da52))
* **deps:** update docker/login-action digest to ec9cdf0 ([#437](https://github.com/open-feature/flagd/issues/437)) ([2650670](https://github.com/open-feature/flagd/commit/2650670d35166e119f9a92613d3aca81523b9faa))
* **deps:** update docker/metadata-action digest to 3343011 ([#438](https://github.com/open-feature/flagd/issues/438)) ([e7ebf32](https://github.com/open-feature/flagd/commit/e7ebf32caf0eae7449e673da0c10998f97ebf781))
* **deps:** update github/codeql-action digest to 32dc499 ([#439](https://github.com/open-feature/flagd/issues/439)) ([f91d91b](https://github.com/open-feature/flagd/commit/f91d91bf020d330f96572c5ee11a210c0c7f4311))
* **deps:** update google-github-actions/release-please-action digest to d3c71f9 ([#406](https://github.com/open-feature/flagd/issues/406)) ([6e1ffb2](https://github.com/open-feature/flagd/commit/6e1ffb27fea5e91014a6991b2afca9a59f89117f))
* disable caching tests in CI ([#442](https://github.com/open-feature/flagd/issues/442)) ([28a35f6](https://github.com/open-feature/flagd/commit/28a35f62d618539362ae83a48f11af08ca2ae245))
* fix race condition on init read ([#409](https://github.com/open-feature/flagd/issues/409)) ([0c9eb23](https://github.com/open-feature/flagd/commit/0c9eb2322df99b4216d40afd1cb3b8873b0c59ff))
* integration test stability ([#432](https://github.com/open-feature/flagd/issues/432)) ([5a6a5d5](https://github.com/open-feature/flagd/commit/5a6a5d5887badd846cffe882c8c22a35b850fa06))
* integration tests ([#312](https://github.com/open-feature/flagd/issues/312)) ([6192ac8](https://github.com/open-feature/flagd/commit/6192ac8820b0f472672ba177b7c5838244b6e277))
* reorder release note sections ([df7bfce](https://github.com/open-feature/flagd/commit/df7bfce85ec7d6abaa987f87341c5af380904b51))
* use -short flag in benchmark tests ([#431](https://github.com/open-feature/flagd/issues/431)) ([e68a6aa](https://github.com/open-feature/flagd/commit/e68a6aadb3dac46676299ab94a34a0bcc39a67af))


### 🐛 Bug Fixes

* **deps:** update kubernetes packages to v0.26.2 ([#450](https://github.com/open-feature/flagd/issues/450)) ([2885227](https://github.com/open-feature/flagd/commit/28852270f34ff81c072337b29aa17f4b6634e9cc))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.2 ([#416](https://github.com/open-feature/flagd/issues/416)) ([feb7f04](https://github.com/open-feature/flagd/commit/feb7f047365263758a63d8dffea936f621a4966d))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.9 ([#427](https://github.com/open-feature/flagd/issues/427)) ([42d2705](https://github.com/open-feature/flagd/commit/42d270558bf9badcff9a9b352fda35491c45aebe))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.29 ([#429](https://github.com/open-feature/flagd/issues/429)) ([b7fae81](https://github.com/open-feature/flagd/commit/b7fae81b89b3a1a0793a688c32569c4284633c6a))
* **deps:** update module github.com/stretchr/testify to v1.8.2 ([#440](https://github.com/open-feature/flagd/issues/440)) ([ab3e674](https://github.com/open-feature/flagd/commit/ab3e6748abc7843c022afeaf7cb11193cdcf59c5))
* **deps:** update module golang.org/x/net to v0.7.0 ([#410](https://github.com/open-feature/flagd/issues/410)) ([c6133b6](https://github.com/open-feature/flagd/commit/c6133b6af61f3d73ae73d318a1a9f44db2540540))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.5 ([#454](https://github.com/open-feature/flagd/issues/454)) ([f907f11](https://github.com/open-feature/flagd/commit/f907f114f23fa2efa2637e254e829e4d53a90b51))
* remove non-error error log from parseFractionalEvaluationData ([#446](https://github.com/open-feature/flagd/issues/446)) ([34aca79](https://github.com/open-feature/flagd/commit/34aca79e6ec9876a6cced0fe49e1ceea34d83696))


### ✨ New Features

* add debug logging for merge behaviour  ([#456](https://github.com/open-feature/flagd/issues/456)) ([dc71e84](https://github.com/open-feature/flagd/commit/dc71e84f0704690b528e7f1c2b56cb4898374fbf))
* add Health and Readiness probes ([#418](https://github.com/open-feature/flagd/issues/418)) ([7f2358c](https://github.com/open-feature/flagd/commit/7f2358ce207527c890f4a2f46ce4b9e8bf697095))
* Add version to startup message ([#430](https://github.com/open-feature/flagd/issues/430)) ([8daf613](https://github.com/open-feature/flagd/commit/8daf613e7e4f4492df0c06e2ef464f4337cadaca))
* introduce flag merge behaviour ([#414](https://github.com/open-feature/flagd/issues/414)) ([524f65e](https://github.com/open-feature/flagd/commit/524f65ea7215466bb4ac24a8d0d5953dd1cfe9a0))
* introduce grpc sync for flagd ([#297](https://github.com/open-feature/flagd/issues/297)) ([33413f2](https://github.com/open-feature/flagd/commit/33413f25882a3f1cf4953da0f18e746bfb69faf4))
* refactor and improve K8s sync provider ([#443](https://github.com/open-feature/flagd/issues/443)) ([4c03bfc](https://github.com/open-feature/flagd/commit/4c03bfc812e7ceabcac0979290bd74d9efc9da15))
* Use OTel to export metrics (metric name changes) ([#419](https://github.com/open-feature/flagd/issues/419)) ([eb3982a](https://github.com/open-feature/flagd/commit/eb3982a1cb72d664022b5cb126b533cf61497001))


### 📚 Documentation

* add .net flagd provider ([73d7840](https://github.com/open-feature/flagd/commit/73d7840c9fdef9c62371c677e02c0d9773c85f95))
* configuration merge docs ([#455](https://github.com/open-feature/flagd/issues/455)) ([6cb66b1](https://github.com/open-feature/flagd/commit/6cb66b14d01b6ee1c270bbdd3e30d4016757eae5))
* documentation for creating a provider ([#413](https://github.com/open-feature/flagd/issues/413)) ([d0c099d](https://github.com/open-feature/flagd/commit/d0c099d9aba3ed4d760a1858381f5e29b6d49a9c))
* updated filepaths for schema store regex ([#344](https://github.com/open-feature/flagd/issues/344)) ([2d0e9d9](https://github.com/open-feature/flagd/commit/2d0e9d956fbc99f2775821cfecdceb2b016d2b78))

## [0.3.7](https://github.com/open-feature/flagd/compare/v0.3.6...v0.3.7) (2023-02-13)


### Bug Fixes

* **deps:** update module golang.org/x/net to v0.6.0 ([#396](https://github.com/open-feature/flagd/issues/396)) ([beb7564](https://github.com/open-feature/flagd/commit/beb756470b1e1d5ef0670b8322b6ed9cb44efa24))
* **deps:** update module google.golang.org/grpc to v1.53.0 ([#388](https://github.com/open-feature/flagd/issues/388)) ([174cd7c](https://github.com/open-feature/flagd/commit/174cd7c70fa5ae2573db2c5972b75786633e2f41))
* error handling of Serve/ServeTLS funcs ([#397](https://github.com/open-feature/flagd/issues/397)) ([8923bf2](https://github.com/open-feature/flagd/commit/8923bf2d407e18b65c188aef9bf7370fc74c3be2))
* fix race in http sync test ([#401](https://github.com/open-feature/flagd/issues/401)) ([1d0c8e1](https://github.com/open-feature/flagd/commit/1d0c8e168b73f7fbd4b27ece733041bbe08261c0))
* sbom artefact name ([#380](https://github.com/open-feature/flagd/issues/380)) ([3daef26](https://github.com/open-feature/flagd/commit/3daef263c43ed63776d604d27f7ae6b993fff143)), closes [#379](https://github.com/open-feature/flagd/issues/379)

## [0.3.6](https://github.com/open-feature/flagd/compare/v0.3.5...v0.3.6) (2023-02-06)


### Bug Fixes

* set ResolveObject reason ([#375](https://github.com/open-feature/flagd/issues/375)) ([dcf199d](https://github.com/open-feature/flagd/commit/dcf199dab9d11b86454028869a54d77a474fc4a6))

## [0.3.5](https://github.com/open-feature/flagd/compare/v0.3.4...v0.3.5) (2023-02-06)


### Features

* flagd image signing ([#338](https://github.com/open-feature/flagd/issues/338)) ([eca6a60](https://github.com/open-feature/flagd/commit/eca6a60967999a303ceef5465f1acc35c83afd6d))
* update in logging to console and Unify case usage, seperators and punctuation for logging ([#322](https://github.com/open-feature/flagd/issues/322)) ([0bdcfd2](https://github.com/open-feature/flagd/commit/0bdcfd2fecc03b15be9fc4b0489431b8fa86aed8))


### Bug Fixes

* **deps:** update module github.com/bufbuild/connect-go to v1.5.1 ([#365](https://github.com/open-feature/flagd/issues/365)) ([e25f452](https://github.com/open-feature/flagd/commit/e25f452906e034e339309270cc8db6dcd58e9973))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.28 ([#342](https://github.com/open-feature/flagd/issues/342)) ([e6df80f](https://github.com/open-feature/flagd/commit/e6df80fd25d3da342e72d2ca0e923d9bf3d3f797))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.2 ([#336](https://github.com/open-feature/flagd/issues/336)) ([836d3cf](https://github.com/open-feature/flagd/commit/836d3cf3c06570d59929c3464e3c8e11c9b5a2fa))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.3 ([#372](https://github.com/open-feature/flagd/issues/372)) ([330ac91](https://github.com/open-feature/flagd/commit/330ac91e375124826b2a7a1a22d0daa18368ab99))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.4 ([#374](https://github.com/open-feature/flagd/issues/374)) ([d90e561](https://github.com/open-feature/flagd/commit/d90e561bfc5b798d13d4ba8f30f523b1053f3748))
* fix unbuffered channel blocking goroutine  ([#358](https://github.com/open-feature/flagd/issues/358)) ([4f1905a](https://github.com/open-feature/flagd/commit/4f1905a9ac6d62b5edb297fba904aac8680c89cf))
* introduced RWMutex to flag state to prevent concurrent r/w of map ([#370](https://github.com/open-feature/flagd/issues/370)) ([93e356b](https://github.com/open-feature/flagd/commit/93e356b4ab0b65c71659bd52d73f618edffc96f5))
* use event.Has func for file change notification handling (increased stability across OS) ([#361](https://github.com/open-feature/flagd/issues/361)) ([09f74b9](https://github.com/open-feature/flagd/commit/09f74b9c5d15622c98da08558cbcd63fe9422754))

## [0.3.4](https://github.com/open-feature/flagd/compare/v0.3.3...v0.3.4) (2023-01-28)


### Bug Fixes

* **deps:** update goreleaser/goreleaser-action action to v4 ([#340](https://github.com/open-feature/flagd/issues/340)) ([b9fcd5c](https://github.com/open-feature/flagd/commit/b9fcd5caa67a61b447c437b651471b4603b2b272))

## [0.3.3](https://github.com/open-feature/flagd/compare/v0.3.2...v0.3.3) (2023-01-28)


### Bug Fixes

* **deps:** update module github.com/bufbuild/connect-go to v1.5.0 ([#326](https://github.com/open-feature/flagd/issues/326)) ([7f332e5](https://github.com/open-feature/flagd/commit/7f332e50ecb1cea19108d1fa2fd79da3d5864bf9))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.26 ([#331](https://github.com/open-feature/flagd/issues/331)) ([be67e5f](https://github.com/open-feature/flagd/commit/be67e5f5bc1fb7351a04ffc4180447a27d57d32a))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.27 ([#335](https://github.com/open-feature/flagd/issues/335)) ([824cf1a](https://github.com/open-feature/flagd/commit/824cf1ab0f2e18826207af16d5328b817c755c8e))
* send datasync on remove fs events ([#339](https://github.com/open-feature/flagd/issues/339)) ([4c9aaac](https://github.com/open-feature/flagd/commit/4c9aaaca77b1c8b16f59434aeb37407fced47ecf))

## [0.3.2](https://github.com/open-feature/flagd/compare/v0.3.1...v0.3.2) (2023-01-26)


### Bug Fixes

* deprecation warning fix ([#317](https://github.com/open-feature/flagd/issues/317)) ([a2630db](https://github.com/open-feature/flagd/commit/a2630dbba151f35cc985d38de9cf25bfee2b76c8))
* **deps:** update kubernetes packages to v0.26.1 ([#267](https://github.com/open-feature/flagd/issues/267)) ([26825f2](https://github.com/open-feature/flagd/commit/26825f288c56df638fd160caa93f926a8c136108))
* **deps:** update module github.com/diegoholiveira/jsonlogic/v3 to v3.2.7 ([#283](https://github.com/open-feature/flagd/issues/283)) ([2ab5a00](https://github.com/open-feature/flagd/commit/2ab5a00fa6f19c7e0fe1a4e36649fae2633ac269))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.24 ([#290](https://github.com/open-feature/flagd/issues/290)) ([38d3eba](https://github.com/open-feature/flagd/commit/38d3ebaffcb1f36a38003273c62c6317f0ee75a3))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.25 ([#324](https://github.com/open-feature/flagd/issues/324)) ([ed1d3aa](https://github.com/open-feature/flagd/commit/ed1d3aaba4ca179a89757a6b1c3f328826e787fc))
* **deps:** update module github.com/spf13/viper to v1.15.0 ([#296](https://github.com/open-feature/flagd/issues/296)) ([d43220b](https://github.com/open-feature/flagd/commit/d43220b2be58e4bce05050c5d1b36788289ae7cc))
* **deps:** update module google.golang.org/grpc to v1.52.1 ([#314](https://github.com/open-feature/flagd/issues/314)) ([ad25388](https://github.com/open-feature/flagd/commit/ad25388461816100e19bda44a8e0077770ea0ee4))
* **deps:** update module google.golang.org/grpc to v1.52.3 ([#325](https://github.com/open-feature/flagd/issues/325)) ([8013ea5](https://github.com/open-feature/flagd/commit/8013ea5c6fa311b337c7ec1b1e8e756080808948))
* Update flagd systemd config to use URI ([#315](https://github.com/open-feature/flagd/issues/315)) ([93a04b4](https://github.com/open-feature/flagd/commit/93a04b46133e9220ec6f23d833c11f195e05c13e))
* update outdated doc link in deprecation warning ([#316](https://github.com/open-feature/flagd/issues/316)) ([19695d2](https://github.com/open-feature/flagd/commit/19695d2715129d6718ca0617b6ec6922ffb79c9b))

## [0.3.1](https://github.com/open-feature/flagd/compare/v0.3.0...v0.3.1) (2023-01-12)


### Features

* file extension detection ([#257](https://github.com/open-feature/flagd/issues/257)) ([ca22541](https://github.com/open-feature/flagd/commit/ca2254117adc163b94d662b3d1fbfd868f788fcb))
* ResolveAll endpoint for bulk evaluation ([#239](https://github.com/open-feature/flagd/issues/239)) ([6437c43](https://github.com/open-feature/flagd/commit/6437c43022b5c94d2fb835a406d85a4e836f2fcf))


### Bug Fixes

* **deps:** update module github.com/bufbuild/connect-go to v1.4.1 ([#268](https://github.com/open-feature/flagd/issues/268)) ([712d7dd](https://github.com/open-feature/flagd/commit/712d7dd4a34980bf9eddad99d926cbdd5d69d624))
* **deps:** update module github.com/mattn/go-colorable to v0.1.13 ([#260](https://github.com/open-feature/flagd/issues/260)) ([5b11504](https://github.com/open-feature/flagd/commit/5b11504cdce50c137540cc79d2db94e70a21338b))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.23 ([#261](https://github.com/open-feature/flagd/issues/261)) ([a1dd3b9](https://github.com/open-feature/flagd/commit/a1dd3b9005374b5527f12b8e138250cacddc71af))
* **deps:** update module github.com/rs/cors to v1.8.3 ([#264](https://github.com/open-feature/flagd/issues/264)) ([0e6f2f3](https://github.com/open-feature/flagd/commit/0e6f2f3e5a77dae7d491eaf1094a65e692bebe5d))
* **deps:** update module github.com/stretchr/testify to v1.8.1 ([#265](https://github.com/open-feature/flagd/issues/265)) ([2ec61c6](https://github.com/open-feature/flagd/commit/2ec61c6bc61c266451b496ff18c3dd9a74173233))
* improve invalid sync URI errror msg ([#252](https://github.com/open-feature/flagd/issues/252)) ([5939870](https://github.com/open-feature/flagd/commit/5939870b8994dbca585c53dd022485090aab2406))
* replace character slice with regex replace ([#250](https://github.com/open-feature/flagd/issues/250)) ([c92d101](https://github.com/open-feature/flagd/commit/c92d1012b0de6af694c3af2fede28053e2572b04))

## [0.3.0](https://github.com/open-feature/flagd/compare/v0.2.7...v0.3.0) (2023-01-06)


### ⚠ BREAKING CHANGES

* consolidated configuration change events into one event ([#241](https://github.com/open-feature/flagd/issues/241))

### Features

* consolidated configuration change events into one event ([#241](https://github.com/open-feature/flagd/issues/241)) ([f9684b8](https://github.com/open-feature/flagd/commit/f9684b858dfef40576e0031654b421a37e8bb1d6))
* support yaml evaluator ([#206](https://github.com/open-feature/flagd/issues/206)) ([2dbace5](https://github.com/open-feature/flagd/commit/2dbace5b6bb8e187a7d44a3d3ec14190c63b3ae0))


### Bug Fixes

* changed eventing configuration mutex to rwmutex and added missing lock ([#220](https://github.com/open-feature/flagd/issues/220)) ([5bbef9e](https://github.com/open-feature/flagd/commit/5bbef9ea4b1960686e58298c2c2e192ca99f072f))
* omitempty targeting field in Flag structure ([#247](https://github.com/open-feature/flagd/issues/247)) ([3f406b5](https://github.com/open-feature/flagd/commit/3f406b53bda8b5beb8b0929da3802a0368c13151))

## [0.2.7](https://github.com/open-feature/flagd/compare/v0.2.5...v0.2.7) (2022-12-02)


### ⚠ BREAKING CHANGES

* start command flag refactor ([#222](https://github.com/open-feature/flagd/issues/222))

### Features

* enable request logging via the --debug flag ([#226](https://github.com/open-feature/flagd/issues/226)) ([11954b5](https://github.com/open-feature/flagd/commit/11954b521cc6197d0dc04b163e66e38d4c288047))
* Resurrected the STATIC flag reason. Documented the caching strategy. ([#224](https://github.com/open-feature/flagd/issues/224)) ([5830592](https://github.com/open-feature/flagd/commit/5830592053c55dc9e55c16854e40c3fc8345d6d1))
* snap ([#211](https://github.com/open-feature/flagd/issues/211)) ([c619844](https://github.com/open-feature/flagd/commit/c61984448d5cdadec62b5cf6f7e24fc5f75a3738))
* start command flag refactor ([#222](https://github.com/open-feature/flagd/issues/222)) ([14474cc](https://github.com/open-feature/flagd/commit/14474ccf65b9b92213e8c792e94c458022484df4))


### Miscellaneous Chores

* release v0.2.6 ([93cfb78](https://github.com/open-feature/flagd/commit/93cfb78d024b436fa7fb17fd41f74d1508bf8b64))
* release v0.2.7 ([4a9f6df](https://github.com/open-feature/flagd/commit/4a9f6df4e472229ff805e9d5d3aa581c7c9c0667))

## [0.2.5](https://github.com/open-feature/flagd/compare/v0.2.4...v0.2.5) (2022-10-20)


### Bug Fixes

* CVE-2022-32149 ([#198](https://github.com/open-feature/flagd/issues/198)) ([11a7b34](https://github.com/open-feature/flagd/commit/11a7b3472ab2bc39bce7c40037e8f83736065163))

## [0.2.4](https://github.com/open-feature/flagd/compare/v0.2.3...v0.2.4) (2022-10-14)


### Bug Fixes

* ApiVersion check fix ([#193](https://github.com/open-feature/flagd/issues/193)) ([3a524d6](https://github.com/open-feature/flagd/commit/3a524d646187355bb224100f436c7b5f35abea3e))

## [0.2.3](https://github.com/open-feature/flagd/compare/v0.2.2...v0.2.3) (2022-10-13)


### Features

* Eventing ([#187](https://github.com/open-feature/flagd/issues/187)) ([3f7fcd2](https://github.com/open-feature/flagd/commit/3f7fcd2f57318fad4e0bf501f202af990d3c5a79))
* fixing informer issues ([#191](https://github.com/open-feature/flagd/issues/191)) ([837b0c6](https://github.com/open-feature/flagd/commit/837b0c673e7e7d4799f100291ca520d22944f22a))
* only fire modify event when FeatureFlagConfiguration Generation field has changed ([#167](https://github.com/open-feature/flagd/issues/167)) ([e2fc7ee](https://github.com/open-feature/flagd/commit/e2fc7ee2570a119923bf95b40b2046dfa4705f20))

## [0.2.2](https://github.com/open-feature/flagd/compare/v0.2.1...v0.2.2) (2022-10-03)


### Bug Fixes

* updated merge functionality ([#182](https://github.com/open-feature/flagd/issues/182)) ([94d7697](https://github.com/open-feature/flagd/commit/94d7697d08a07cede4a548ef998792d00f8954a0))

## [0.2.1](https://github.com/open-feature/flagd/compare/v0.2.0...v0.2.1) (2022-09-27)


### Bug Fixes

* updated tcp listener ([#174](https://github.com/open-feature/flagd/issues/174)) ([b750ed1](https://github.com/open-feature/flagd/commit/b750ed1268b5e6efe779a34e764bad3e781f8e93))

## [0.2.0](https://github.com/open-feature/flagd/compare/v0.1.1...v0.2.0) (2022-09-26)


### ⚠ BREAKING CHANGES

* Updated service to use connect (#163)

### Features

* Updated service to use connect ([#163](https://github.com/open-feature/flagd/issues/163)) ([828d5c4](https://github.com/open-feature/flagd/commit/828d5c4c11157f5b7a77f5041806ba2523186764))


### Bug Fixes

* checkout release tag before running container and binary releases ([#171](https://github.com/open-feature/flagd/issues/171)) ([50fe46f](https://github.com/open-feature/flagd/commit/50fe46fbbf120a0657c1df35b370cdc9051d0f31))

## [0.1.1](https://github.com/open-feature/flagd/compare/v0.1.0...v0.1.1) (2022-09-23)


### Bug Fixes

* bubbles up unclean error exits ([#170](https://github.com/open-feature/flagd/issues/170)) ([9f7db02](https://github.com/open-feature/flagd/commit/9f7db0259d2d24cb880eeddaebd3b8f48758248a))
* upgrade package containing vulnerability ([#162](https://github.com/open-feature/flagd/issues/162)) ([82278c7](https://github.com/open-feature/flagd/commit/82278c7cf08cc6b50f49ab500caf6f9003fc0823))
