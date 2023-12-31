package main

import (
    "fmt"
    "html/template"
    "net/http"
    "github.com/ryankenney-dev/advent-of-code-2023/day1"
    "github.com/ryankenney-dev/advent-of-code-2023/day2_part1"
    "github.com/ryankenney-dev/advent-of-code-2023/day2_part2"
    "github.com/ryankenney-dev/advent-of-code-2023/day3"
    "github.com/ryankenney-dev/advent-of-code-2023/day3v2"
    "github.com/ryankenney-dev/advent-of-code-2023/day4_part1"
    "github.com/ryankenney-dev/advent-of-code-2023/day4_part2"
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
                <style>
                    body {
                        background-color: #121212;
                        color: #ffffff;
                        font-family: 'Arial', sans-serif;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        height: 100vh;
                        margin: 0;
                    }

                    #container {
                        text-align: center;
                        width: 75%;
                    }

                    textarea, select, button {
                        margin-bottom: 10px;
                        background-color: #333333;
                        color: white;
                        border: none;
                        padding: 10px;
                        border-radius: 4px;
                    }

                    textarea {
                        width: 100%;
                        box-sizing: border-box; /* Ensures padding doesn't add to the width */
                    }

                    #result {
                        margin-top: 20px;
                        border: 1px solid #ffffff;
                        padding: 10px;
                    }
                </style>
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
                <div id="container">
                    <textarea id="text" rows="10"></textarea><br>
                    <select id="computation">
                        {{range $key, $value := .}}
                            <option value="{{$key}}">{{$value.Title}}</option>
                        {{end}}
                    </select>
                    <button onclick="sendRequest()">Submit</button>
                    <div id="result"></div>
                </div>
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
    "day1part1": day1.Day1Part1{},
    "day1part2": day1.Day1Part2{},
    "day2part1": day2_part1.Day2Part1{},
    "day2part2": day2_part2.Day2Part2{},
    "day3part1": day3.Day3Part1{},
    "day3part2": day3.Day3Part2{},
    "day3v2part1": day3v2.Day3V2Part1{},
    "day4part1": day4_part1.Day4Part1{},
    "day4part2": day4_part2.Day4Part2{},
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

