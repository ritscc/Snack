use std::env::current_dir;

fn main() -> Result<(), std::io::Error> {
    println!("cargo:rerun-if-changed=build.rs");
    let current_dir = current_dir().unwrap();
    let parent_dir = current_dir.parent().unwrap();
    let proto_root_path = parent_dir.parent().unwrap().to_str().unwrap();
    let proto_path = [
        "authentication/v1/authentication.proto",
        "channel/v1/channel.proto",
        "event/v1/event.proto",
        "message/v1/message.proto",
        "stamp/v1/stamp.proto",
        "user/v1/user.proto",
        "user_group/v1/user_group.proto",
    ];
    let protos = proto_path.map(|path| proto_root_path.to_string() + "/protobuf/" + path);
    let build = tonic_build::configure()
        .out_dir("src/api")
        .compile_well_known_types(true)
        .build_server(true)
        .build_client(true)
        .emit_rerun_if_changed(true)
        .compile(&protos, &[proto_root_path]);

    std::process::Command::new("cargo")
        .arg("fmt")
        .status()
        .expect("Failed to run 'cargo fmt'");

    build
}
