import React from "react"
import Comment from "./comment"

const CommentList = ({comments}) => (
  <div>
    {comments.map((c) => <Comment register={c.register} text={c.text} />)}
  </div>
)

export default CommentList
