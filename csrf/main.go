// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/surma/httptools"
	"github.com/udacity/ud897-client-server-communication/utils"
)

//go:generate rice embed-go
var box *rice.Box

func main() {
	box = rice.MustFindBox("assets")
	os.Mkdir("evil", os.FileMode(0755))
	log.Printf("Created a folder called \"evil\", all its contents")
	log.Printf("will be served on evil.127.0.0.1.xip.io:8080")
	log.Printf("Running bank website on bank.127.0.0.1.xip.io:8080")

	bank := httptools.NewRegexpSwitch(map[string]http.Handler{
		"/": http.RedirectHandler("/balance", http.StatusTemporaryRedirect),
		"/transfer": httptools.MethodSwitch{
			"POST": httptools.List{
				httptools.SilentHandlerFunc(checkLogin),
				http.HandlerFunc(transfer),
			},
		},
		"/login": httptools.MethodSwitch{
			"GET":  http.HandlerFunc(showLoginForm),
			"POST": http.HandlerFunc(login),
		},
		"/logout": http.HandlerFunc(logout),
		"/balance": httptools.MethodSwitch{
			"GET": httptools.List{
				httptools.SilentHandlerFunc(checkLogin),
				http.HandlerFunc(showBalance),
			},
		},
	})

	evil := http.FileServer(http.Dir("evil"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname := strings.Split(r.Host, ":")[0]
		switch {
		case strings.HasPrefix(hostname, "evil."):
			evil.ServeHTTP(w, r)
		case strings.HasPrefix(hostname, "bank."):
			bank.ServeHTTP(w, r)
		}
	}))
	http.ListenAndServe(":8080", nil)
}

type TransferAction struct {
	Amount    string
	Recipient string
	Date      string
}

type Account struct {
	Transfers []TransferAction
	*sync.Mutex
}

var (
	acc = &Account{
		Transfers: []TransferAction{
			{
				Amount:    "1337",
				Recipient: "Udacity",
				Date:      time.Now().String(),
			},
		},
		Mutex: &sync.Mutex{},
	}
)

func checkLogin(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("SESSION_ID")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	if c.Value != "totally_secret" {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
}

var (
	secret = "setyourcorsheader"
)

func transfer(w http.ResponseWriter, r *http.Request) {
	t := TransferAction{
		Amount:    r.FormValue("amount"),
		Recipient: r.FormValue("recipient"),
	}
	t.Date = time.Now().String()
	acc.Lock()
	defer acc.Unlock()
	acc.Transfers = append(acc.Transfers, t)

	refUrl, err := url.Parse(r.Referer())
	if err == nil && strings.HasPrefix(refUrl.Host, "evil.") && t.Recipient == "Evil Corp" {
		log.Printf("You made Evil Corp(tm) rich!")
		log.Printf("The oken is >> %s <<", secret)
	}
	http.Redirect(w, r, "/balance", http.StatusSeeOther)
}

func showBalance(w http.ResponseWriter, r *http.Request) {
	err := utils.ExecuteTemplateInBox(w, box, "balance.html", acc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error executing template: %s", err)
		return
	}
}

func showLoginForm(w http.ResponseWriter, r *http.Request) {
	err := utils.ExecuteTemplateInBox(w, box, "login.html", acc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error executing template: %s", err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("password") == "super secret password" {
		http.SetCookie(w, &http.Cookie{
			Name:  "SESSION_ID",
			Value: "totally_secret",
		})
		http.Redirect(w, r, "/balance", http.StatusSeeOther)
		return
	}
	http.Error(w, "Wrong password", http.StatusForbidden)
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "SESSION_ID",
		Expires: time.Unix(0, 0),
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
