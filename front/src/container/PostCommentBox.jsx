import { connect } from "react-redux"
import { OnPostClick, OnNameInputChange, OnTextInputChange} from "../action/comment"
import CommentBox from "../components/CommentBox"

const mapStateToProps = (state) => state.comment

const mapDispatchToProps = (dispatch) => {
  return {
    onPostClick: ({ register, text }) => {
      fetch("/comment", {
        method: "POST",
        body: JSON.stringify({ register: register, text: text }),
      }).then(() => dispatch(OnPostClick))
        .catch((err) => console.error(err))
    },
    onNameChange: (event) => {
      dispatch(OnNameInputChange(event.target.value))
    },
    onTextChange: (event) => {
      dispatch(OnTextInputChange(event.target.value))
    }
  }
}

const PostCommentBox = connect(mapStateToProps, mapDispatchToProps)(CommentBox)

export default PostCommentBox
