import React from "react"
import ReactDOM from "react-dom"

import CommentList from "./components/CommentList"
import CommentBox from "./components/CommentBox"

const comments = [
  {
    register: "User1",
    text: "hogehgoe",
  },
  {
    register: "User2",
    text: "hogehgoe",
  },
  {
    register: "User3",
    text: "hogehgoe",
  },
]

const App = () =>
  <div>
    <h1>なんちゃってチャット</h1>
    <CommentList comments={comments} />
    <CommentBox />
  </div>

ReactDOM.render(<App />, document.querySelector("#app"))
