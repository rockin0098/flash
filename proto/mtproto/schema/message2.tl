// test
rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
message2 msg_id:long seqno:int bytes:int body:Object = message2; // parsed manually
msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
msg_copy#e06046b2 orig_message:Message = MessageCopy; // parsed manually, not used - use msg_container
gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually