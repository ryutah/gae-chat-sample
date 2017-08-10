import { connect } from "react-redux"
import CommentList from "../components/CommentList"

const mapStateToProps = (state) => state.comment

const AppendableCommentList = connect(mapStateToProps)(CommentList)

export default AppendableCommentList
