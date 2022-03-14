package main

// Range over endpoint attributes
// for property, val := range meta.Value.Properties {
// 	propDescription := val.Value.Description
// 	fmt.Println("attribute:", property)
// 	fmt.Println("placeholder:", propDescription)
// }

const webHTMLServiceTemplate = `
{{ $service := .service -}}
<html>
  <body>
    <div id="{{ $service.Name }}">
      <form id="{{ untitle .endpoint }}" onsubmit="{{ $service.Name }}{{ .endpoint }}()">
        <div>
            <label for="service"><b>{{ $service.Name }}</b></label>
            <input type="hidden" name="service" id="service" value="{{ $service.Name }}">
        </div>
        <div>
            <label for="endpoint"><b>{{ .endpoint }}</b></label>
            <input type="hidden" name="endpoint" id="endpoint" value="{{ .endpoint }}">
        </div>
		<i>{{ .epdesc }}</i>
		</br>
		</br>
		<label for="token">token </label>
        <div>
            <input name="token" id="token" placeholder="token">
        </div>
		</br>
		{{- range $property, $val := .properties }}
		{{- if not (eq $val.Value.Type "object") }}
		<label for="{{ $property }}">{{ $property }} </label>
        <div>
            <input name="{{ $property }}" id="{{ $property }}" placeholder="{{ $val.Value.Description }}">
        </div>
		{{- end }}
		{{- if eq $val.Value.Type "object" }}
		<label for="{{ $property }}">{{ $property }} </label>
        <div>
            <textarea rows=5 cols=30 name="{{ $property }}" id="{{ $property }}" placeholder="{{ $val.Value.Description }}">{}</textarea>
        </div>
		{{- end }}
		{{- end }}
        <button>Submit</button>
      </form>
    </div>
    <div id="response"></div>
  </body>
  <script type="module" src="{{ untitle .endpoint }}.js"></script>
</html>
`

const webJSServiceTemplate = `
{{ $service := .service -}}
import Client from '../../client/index';

function {{ $service.Name }}{{ .endpoint }}() {
	var form = form.elements["{{ untitle .endpoint }}"].value;
	var token = form.elements["token"].value;
	var service = form.elements["service"].value;
	var endpoint = form.elements["endpoint"].value;
	{{- range $property, $val := .properties }}
	var {{ $property }} = form.elements["{{ $property }}"].value;
	{{- end }}
	var obj = new Object();
	{{- range $property, $val := .properties }}
	obj.{{ $property }} = {{ $property }};
	{{ end }}
	var request = JSON.stringify(obj);

	var m3o = new Client(token);

	m3o.call(service, endpoint, request, function(response) {
		document.getElementById("response").innerText = response;
	});
}
`
