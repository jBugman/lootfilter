import util.{quoted}

pub type BaseType =
  String

pub fn to_string(base_type: BaseType) -> String {
  quoted(base_type)
}
