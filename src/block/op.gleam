pub type Op {
  Eq
  NotEq
  LT
  LTE
  GT
  GTE
  ExactMatch
}

pub fn to_string(op: Op) -> String {
  case op {
    Eq -> "="
    NotEq -> "!="
    LT -> "<"
    LTE -> "<="
    GT -> ">"
    GTE -> ">="
    ExactMatch -> "=="
  }
}
