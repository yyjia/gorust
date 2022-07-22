extern crate libc;

extern "C" {
    fn goGetNonce(address: *const libc::c_char, block_num: libc::c_int) -> libc::c_int;
}

#[no_mangle]
pub extern "C" fn execute(address: *const libc::c_char, block_num: libc::c_int) -> libc::c_int {
    unsafe {
        let nonce = goGetNonce(address, block_num);
        println!(
            "call rust execute({:?}, {})-> {}",
            address.as_ref().unwrap(),
            block_num,
            nonce
        );
        nonce
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::ffi::CString;

    #[test]
    fn it_works() {
        assert_eq!(
            execute(CString::new("abcdefg").unwrap().into_raw(), 1234),
            1000
        );
        assert_eq!(
            execute(CString::new("xxxxxxx").unwrap().into_raw(), 1234),
            10
        );
    }
}
