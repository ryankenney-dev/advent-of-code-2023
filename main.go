package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/compute", compute)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("form").Parse(`
        <html>
            <head>
                <script>
                    function sendRequest() {
                        var xhr = new XMLHttpRequest();
                        xhr.open("POST", "/compute", true);
                        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
                        xhr.onreadystatechange = function() {
                            if (this.readyState == 4 && this.status == 200) {
                                document.getElementById("result").innerText = this.responseText;
                            }
                        };
                        var text = document.getElementById("text").value;
                        var computation = document.getElementById("computation").value;
                        xhr.send("text=" + encodeURIComponent(text) + "&computation=" + encodeURIComponent(computation));
                    }
                </script>
            </head>
            <body>
                <textarea id="text" rows="10" cols="30"></textarea><br>
                <select id="computation">
                    {{range $key, $value := .}}
                        <option value="{{$key}}">{{$value.Title}}</option>
                    {{end}}
                </select>
                <button onclick="sendRequest()">Submit</button>
                <div id="result"></div>
            </body>
        </html>
    `))

    // Create a data structure for template
    data := make(map[string]struct{ Title string })
    for key, alg := range algorithms {
        data[key] = struct{ Title string }{Title: alg.Title()}
    }

    tmpl.Execute(w, data)
}

var algorithms = map[string]Algorithm{
    "lines": LinesAlgorithm{},
    "characters": CharactersAlgorithm{},
    "day1part1": Day1Part1{},
    "day1part2": Day1Part2{},
    "day2part1": Day2Part1{},
    "day2part2": Day2Part2{},
}

func compute(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    r.ParseForm()
    text := r.FormValue("text")
    computationType := r.FormValue("computation")

    var result string
    if algorithm, ok := algorithms[computationType]; ok {
        result = algorithm.Compute(text)
    } else {
        result = "Invalid computation type"
    }

    fmt.Fprintf(w, result)
}

