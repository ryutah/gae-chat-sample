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
      const newState = Object.assign({}, state)
      newState.comments.push({
        register: state.register,
        text: state.text,
      })
      newState.commentsSize = newState.comments.length
      return newState
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
    default: {
      return state
    }
  }
}
