// +build !js

// Package websocket implements the RFC 6455 WebSocket protocol.
//
// https://tools.ietf.org/html/rfc6455
//
// Use Dial to dial a WebSocket server and Accept to accept a WebSocket client.
// Conn represents the resulting WebSocket connection.
//
// The examples are the best way to understand how to correctly use the library.
//
// The wsjson and wspb subpackages contain helpers for JSON and Protobuf messages.
//
// See https://nhooyr.io/websocket for further information.
//
// Wasm
//
// The client side supports compiling to Wasm.
// It wraps the WebSocket browser API.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket
//
// Some important caveats to be aware of:
//
//  - Conn.Ping is no-op
//  - HTTPClient, HTTPHeader and CompressionOptions in DialOptions are no-op
//  - *http.Response from Dial is &http.Response{} on success
//
package websocket // import "nhooyr.io/websocket"
