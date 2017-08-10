export const TypeOnPostClick       = "ON_POST_CLICK"
export const TypeOnNameInputChange = "ON_NAME_INPUT_CHANGE"
export const TypeOnTextInputChange = "ON_TEXT_INPUT_CHANGE"
export const TypeFetchComment      = "FETCH_COMMENT"

export const OnPostClick = {
  type: TypeOnPostClick,
}

export const OnNameInputChange = (text) => ({
  type: TypeOnNameInputChange,
  text,
})

export const OnTextInputChange = (text) => ({
  type: TypeOnTextInputChange,
  text,
})

export const FetchComment = (comments) => ({
  type: TypeFetchComment,
  comments,
})
