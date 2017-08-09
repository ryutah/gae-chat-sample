export const TypeOnPostClick       = "ON_POST_CLICK"
export const TypeOnNameInputChange = "ON_NAME_INPUT_CHANGE"
export const TypeOnTextInputChange = "ON_TEXT_INPUT_CHANGE"

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
