package main
//
//// 0 OMIT
//// ReaderOverflow returns more than len(b) bytes
//type ReaderOverflow struct{}
//func (r *ReaderOverflow) Read(p []byte) (int, error) {
//	return len(p) + 1, nil
//}
//// END 0 OMIT
//
//// 1 OMIT
//// ReaderMemoryDevourer always returns 1, nil
//type ReaderMemoryDevourer struct{}
//func (r *ReaderMemoryDevourer) Read(p []byte) (int, error) {
//	return 1, nil
//}
//// END 1 OMIT
//
//// 2 OMIT
//// ReaderNegative always returns -1, nil
//type ReaderNegative struct{}
//func (r *ReaderNegative) Read(p []byte) (int, error) {
//	return -1, nil
//}
//// END 2 OMIT
//
//// 3 OMIT
//// ReaderInfinite always returns 0, nil
//type ReaderInfinite struct{}
//func (r *ReaderInfinite) Read(p []byte) (int, error) {
//	return 0, nil
//}
//// END 3 OMIT