import gleam/int
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/string

pub type Conditions {
  Conditions(bases: List(String), min_ilvl: Option(Int), custom: List(String))
}

pub const six_link = Conditions(
  bases: [],
  min_ilvl: None,
  custom: ["Rarity < Unique", "LinkedSockets = 6"],
)

pub fn render(cond: Conditions) -> List(String) {
  [bases(cond.bases), min_ilvl(cond.min_ilvl)]
  |> option.values
  |> list.append(cond.custom)
}

fn bases(xs: List(String)) -> Option(String) {
  case xs {
    [] -> None
    xs ->
      xs
      |> list.map(fn(s) { "\"" <> s <> "\"" })
      |> string.join(" ")
      |> string.append("BaseType ", _)
      |> Some
  }
}

fn min_ilvl(ilvl: Option(Int)) -> Option(String) {
  option.map(ilvl, fn(x) { "ItemLevel >= " <> int.to_string(x) })
}
