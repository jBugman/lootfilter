import gleeunit/should

import poe/conditions/base_types.{BaseType, BaseTypes}

pub fn base_type_test() {
  BaseType("Short Sword")
  |> base_types.to_string
  |> should.equal("BaseType \"Short Sword\"")
}

pub fn base_types_test() {
  BaseTypes(["Short Sword", "Short Bow"])
  |> base_types.to_string
  |> should.equal("BaseType \"Short Sword\" \"Short Bow\"")
}
