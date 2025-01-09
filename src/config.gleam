import gleam/option.{None, Some}
import gleam/string_tree
import tom

import actions/rgba.{RGBA}
import block
import conditions
import decorations.{Decorations}
import domain
import rarity
import styles

pub fn parse(src: String) -> String {
  let assert Ok(src) = tom.parse(src)

  // TODO: array example
  let _ = unwrap_string
  // let assert Ok(chancing_bases) = tom.get_array(cfg, ["chancing_bases"])
  // chancing_bases
  // |> list.map(unwrap_string)
  // |> string.join("\n")
  //
  // tom.get_string(src, ["global_styles", "background_color"])

  let global_styles = styles.parse(src)

  let tree = string_tree.new()

  let six_links = {
    let color = Some(RGBA("184 218 242"))
    block.Show(
      conditions.six_link,
      Decorations(
        rarity: rarity.Epic,
        font_size: Some(42),
        text: color,
        border: color,
        background: None,
        sound: domain.tink_sound,
      ),
    )
  }
  six_links
  |> block.render(global_styles)
  |> string_tree.append(tree, _)
  |> string_tree.to_string
}

fn unwrap_string(value: tom.Toml) -> String {
  let assert tom.String(s) = value
  s
}
