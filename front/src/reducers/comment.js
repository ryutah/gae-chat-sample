import * as commentAction from "../action/comment"

const defaultState = {
  register: "",
  text: "",
  commentsSize: 0,
  comments: [],
}

export default (state = defaultState, action) => {
  switch (action.type) {
    case commentAction.TypeOnPostClick: {
      return state
    }
    case commentAction.TypeOnNameInputChange: {
      const newState = Object.assign({}, state)
      newState.register = action.text
      return newState
    }
    case commentAction.TypeOnTextInputChange: {
      const newState = Object.assign({}, state)
      newState.text = action.text
      return newState
    }
    case commentAction.TypeFetchComment: {
      const newState = Object.assign({}, state)
      newState.comments = action.comments.comments
      newState.commentsSize = action.comments.comments.length
      return newState
    }
    default: {
      return state
    }
  }
}
