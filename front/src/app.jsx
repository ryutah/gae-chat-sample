import React from "react"
import { render } from "react-dom"
import { Provider } from "react-redux"
import { createStore } from "redux"
import reducer from "./reducers"
import AppendableCommentList from "./container/AppendableCommentList"
import PostCommentBox from "./container/PostCommentBox"
import * as commentAction from "./action/comment"

let store = createStore(reducer)

const App = () => (
  <div>
    <h1>なんちゃってチャット</h1>
    <AppendableCommentList />
    <PostCommentBox />
  </div>
)

render(
  <Provider store={store}>
    <App />
  </Provider>
  , document.querySelector("#app")
)

setInterval(() => {
  fetch("/comment")
    .then((resp) => resp.json())
    .then((json) => store.dispatch(commentAction.FetchComment(json)))
    .catch((err) => console.error(err))
}, 1000)
