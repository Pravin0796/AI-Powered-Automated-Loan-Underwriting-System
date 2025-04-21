import { UserServiceClientImpl } from "../proto/user"; // ✅ Correct
import { LoanServiceClientImpl } from "../proto/loan"; // ✅ Correct
// ✅ This is the generated service client class from ts-proto

import { GrpcWebImpl } from '../proto/user';


// ✅ This is the generated gRPC transport class

const transport = new GrpcWebImpl('http://localhost:9090', {
    // options (optional)
    // fetch: customFetch,
    // metadata: customHeaders,
    // debug: true,
});
export const userClient = new UserServiceClientImpl(transport);
export const loanClient = new LoanServiceClientImpl(transport);
// ✅ This creates the client with the correct transport
