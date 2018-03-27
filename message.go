package main

import (
  "github.com/asticode/go-astilectron"
  "./bootstrap"
  "github.com/asticode/go-astilog"
  "github.com/pkg/errors"
  "net/url"
  "net/http"
  "strings"
  "encoding/json"
)

const (
  ClientId                 = "31dbd4e343e060a20bdf"
  ClientSecret             = "c0a20c8b3bd912998779ad639077442e6b8b3ba9"
  RedirectUri              = "http://localhost"
  AuthorizedCodeURL        = "https://github.com/login/oauth/authorize"
  AuthorizedTokenURL       = "https://github.com/login/oauth/access_token"
)


// handleMessages handles messages
func handleMessages(a *astilectron.Astilectron, _ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
  switch m.Name {
  case "login":

    requestCodeUrl := url.Values{}
    requestCodeUrl.Set("response_type", "code")
    requestCodeUrl.Add("client_id", ClientId)
    requestCodeUrl.Add("client_secret", ClientSecret)
    requestCodeUrl.Add("redirect_uri", RedirectUri)

    var ow, _ = a.NewWindow(AuthorizedCodeURL + "?" + requestCodeUrl.Encode(), &astilectron.WindowOptions{
      Center: astilectron.PtrBool(false),
      Height: astilectron.PtrInt(600),
      Width:  astilectron.PtrInt(600),
    })

    ow.On(astilectron.EventNameWindowEventDidGetRedirectRequest, func(e astilectron.Event) (deleteListener bool) {

      // when receive a code parameter, then post api to get token and send token to front-end if successly get token

      result := &struct {
        AccessToken   string  `json:"access_token"`
        TokenType     string  `json:"token_type,omitempty"`
        Scope         string  `json:"scope,omitempty"`
      }{}

      client := &http.Client{}

      requestTokenUrl := url.Values{}
      requestTokenUrl.Set("client_id", ClientId)
      requestTokenUrl.Add("client_secret", ClientSecret)
      requestTokenUrl.Add("code", getURLParam(e.URLNew, "code"))
      requestTokenUrl.Add("redirect_uri", RedirectUri)

      req, _ := http.NewRequest("POST", AuthorizedTokenURL,
        strings.NewReader(requestTokenUrl.Encode()))

      req.Header.Add("Accept", "application/json")
      req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

      resp, _ := client.Do(req)

      json.NewDecoder(resp.Body).Decode(result)

      resp.Body.Close()

      message :=  &bootstrap.MessageOut{
        Name: e.Name,
        Payload: struct {
          AccessToken   string  `json:"access_token"`
          TokenType     string  `json:"token_type,omitempty"`
          Scope         string  `json:"scope,omitempty"`
        }{
          AccessToken: result.AccessToken,
          TokenType: result.TokenType,
          Scope: result.Scope,
        },
      }

      w.SendMessage(message, func(m *astilectron.EventMessage) {
        // Unmarshal
        var s string
        m.Unmarshal(&s)

        // Process message
        astilog.Debugf("received %s", s)
        ow.Close()
      })

      return
    })

    ow.On(astilectron.EventNameWindowEventWillNavigate, func(e astilectron.Event) (deleteListener bool) {

      // as what the above code do

      result := &struct {
        AccessToken   string  `json:"access_token"`
        TokenType     string  `json:"token_type,omitempty"`
        Scope         string  `json:"scope,omitempty"`
      }{}

      client := &http.Client{}

      requestTokenUrl := url.Values{}
      requestTokenUrl.Set("client_id", ClientId)
      requestTokenUrl.Add("client_secret", ClientSecret)
      requestTokenUrl.Add("code", getURLParam(e.URL, "code"))
      requestTokenUrl.Add("redirect_uri", RedirectUri)

      req, _ := http.NewRequest("POST", AuthorizedTokenURL,
        strings.NewReader(requestTokenUrl.Encode()))

      req.Header.Add("Accept", "application/json")
      req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

      resp, _ := client.Do(req)

      json.NewDecoder(resp.Body).Decode(result)

      resp.Body.Close()

      message :=  &bootstrap.MessageOut{
        Name: e.Name,
        Payload: struct {
          AccessToken   string  `json:"access_token"`
          TokenType     string  `json:"token_type,omitempty"`
          Scope         string  `json:"scope,omitempty"`
        }{
          AccessToken: result.AccessToken,
          TokenType: result.TokenType,
          Scope: result.Scope,
        },
      }

      w.SendMessage(message, func(m *astilectron.EventMessage) {
        // Unmarshal
        var s string
        m.Unmarshal(&s)

        // Process message
        astilog.Debugf("received %s", s)
        ow.Close()
      })

      return
    })

    ow.Create()
  }
  return
}


// get url parameter
func getURLParam(targetURL, param string) string {
  query, err := url.Parse(targetURL)
  if err != nil {
    astilog.Error(errors.Wrap(err, "parsing URL failed"))
  }
  value, err := url.ParseQuery(query.RawQuery)
  if err != nil {
    astilog.Error(errors.Wrap(err, "parsing URLQuery failed"))
  }
  return value.Get(param)
}