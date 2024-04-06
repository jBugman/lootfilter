import argv
import config
import erlang_ffi
import filepath
import fs
import gleam/io
import gleam/list
import gleam/result
import gleam/string

pub fn main() {
  case argv.load().arguments {
    ["--debug", path_in] ->
      case debug(path_in) {
        Error(err) -> io.println_error(err)
        Ok(s) -> io.println(s)
      }
    [path_in] ->
      case run(path_in) {
        Error(err) -> io.println_error(err)
        Ok(_) -> io.println("done")
      }
    _ -> io.println("usage: app [input_file]")
  }
}

fn debug(src_path: String) -> Result(String, String) {
  use src <- result.try(fs.read_text(src_path))
  let filter = config.parse(src)
  Ok(filter)
}

fn run(src_path: String) -> Result(Nil, String) {
  use src <- result.try(fs.read_text(src_path))
  use out_path <- result.try(filter_file_path(src_path))
  io.println("writing to " <> out_path)
  let filter = config.parse(src)
  fs.write_text(filter, out_path)
}

const fixed_output_path = ["Documents", "My Games", "Path of Exile"]

fn filter_file_path(src_path: String) -> Result(String, String) {
  let name =
    filepath.base_name(src_path)
    |> filepath.strip_extension
    |> string.append(".filter")
  use homedir <- result.try(erlang_ffi.homedir())
  list.fold(fixed_output_path, homedir, filepath.join)
  |> filepath.join(name)
  |> Ok
}
