pub type Size {
  Large
  Medium
  Small
}

pub fn to_string(size: Size) -> String {
  case size {
    Large -> "0"
    Medium -> "1"
    Small -> "2"
  }
}
