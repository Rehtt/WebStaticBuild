/*此文件为自动生成文件，不要直接做出修改。如果要修改应该修改源文件再执行生成。*/

package web

type RouterInfo struct {
	Data        []byte
	ContentType string
	Gzip        bool
}

func GetRouter(url string) (RouterInfo, bool) {
	r, ok := router[url]
	return r, ok
}

var router = map[string]RouterInfo{
	"./web/favicon.ico": {
		Data:        res82ae931a94fac46b64ff2296b11c90df,
		ContentType: "",
		Gzip:        true,
	},
	"./web/index.html": {
		Data:        resd1e0fb4cc20d34f9ccee9a491faa5b3a,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/app.b7cb370e.css": {
		Data:        rescc20c35e04c4bacd2b612c10ca393a63,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/chunk-238c903c.3c7f5ad9.css": {
		Data:        res347cc58570307e436e159fb7083feffd,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/chunk-26600eb3.5cd9884a.css": {
		Data:        res56b606b2bf43abb3da9dd4e011bf9544,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/chunk-5088bd5d.9a9361c6.css": {
		Data:        res2682caafdba942c36a596a9d0374afd5,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/chunk-ab3ed656.94702ff7.css": {
		Data:        resf254db167b1c15827b53eb3d5c659247,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/chunk-elementUI.68c70ad5.css": {
		Data:        res9d3f8ceb4fbdc40b8baefc3e7635755b,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/css/chunk-libs.3dfb7769.css": {
		Data:        res4b041366a741eaa8668ac97f37bf2e55,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/fonts/element-icons.535877f5.woff": {
		Data:        res3084852792c7ad7d789fcd94a30af33c,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/fonts/element-icons.732389de.ttf": {
		Data:        res7508d493339fcd6c5fd8ab59ef19d49c,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/img/404.a57b6f31.png": {
		Data:        res99aad363db5d6b1df6fdfe22ba4454f7,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/img/404_cloud.0f4bc32b.png": {
		Data:        resc144412bea553287834bb5350f80699a,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/app.2643b543.js": {
		Data:        res5942b247a3f315c5dd496c4c971ca07d,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-238c903c.263e5dc6.js": {
		Data:        res2880d7aef9360de30aa2004be1eef704,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-26600eb3.6f991c49.js": {
		Data:        resd5c71a4437c5caa15804047e94d0f073,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d0c8bf7.aeed24dc.js": {
		Data:        rese56b0f6ad47a4599a7c7bf129327509d,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d0cfaef.188fdb2b.js": {
		Data:        res6e66e7d941b62574f9ea3d34ec970e8f,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d0d0f79.3f658b24.js": {
		Data:        resf8e434d69e649d64c426e0e1d311b0f8,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d0e4b0c.8f10f391.js": {
		Data:        res82319998ca29d6ac523e67ab51ea37a0,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d0e4e1f.40a28c3f.js": {
		Data:        res73bd2a4f0bb47f43f7b9c0f1de7aff41,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d0e944c.fcb11446.js": {
		Data:        resf87f7b5daa3a7b3f3aa371111a7a0e51,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d2104c6.e24ab3c9.js": {
		Data:        res50d6f91f112030dab33b025cc231ed90,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d226cab.1e2799d5.js": {
		Data:        res62f5d36332459f2601fe5bb473d98f06,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-2d229205.3aa01307.js": {
		Data:        res117742f72000a7b81fe273879859a3d3,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-5088bd5d.355610eb.js": {
		Data:        resb4ba22276558e3d63da05fdf4441c9ba,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-ab3ed656.e2857743.js": {
		Data:        resd31c93e91b77bc212bbbbea21ee27b69,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-elementUI.cb459a4a.js": {
		Data:        resab20ca00f329fbce0829407e9b091584,
		ContentType: "",
		Gzip:        true,
	},
	"./web/static/js/chunk-libs.59bbdc99.js": {
		Data:        res4df68196629cb577523d5c381abfbeb8,
		ContentType: "",
		Gzip:        true,
	},
}
