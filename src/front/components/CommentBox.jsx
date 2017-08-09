import React from "react"

const CommentBox = ({register, text, onPostClick, onNameChange, onTextChange}) => (
  <div>
    <input
      type="text"
      placeholder="名前"
      value={register}
      onChange={onNameChange}
    />
    <textarea
      placeholder="投稿内容"
      value={text}
      onChange={onTextChange}
    />
    <button onClick={onPostClick}>送信</button>
  </div>
)

export default CommentBox
