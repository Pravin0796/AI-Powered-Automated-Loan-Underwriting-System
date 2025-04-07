// import { grpc } from "@improbable-eng/grpc-web";
import { UserServiceClient } from "../proto/UserServiceClientPb";
import * as proto from '../proto/user_pb.d';

const RegisterRequest = proto.RegisterRequest;
const LoginRequest = proto.LoginRequest;

const client = new UserServiceClient("http://localhost:50051");

export const registerUser = async (data: { full_name: string; email: string; password: string; phone: string; date_of_birth: string; address: string }) => {
    return new Promise((resolve, reject) => {
        const request = new RegisterRequest();
        request.setFullName(data.full_name);
        request.setEmail(data.email);
        request.setPassword(data.password);
        request.setPhone(data.phone);
        request.setDateOfBirth(data.date_of_birth);
        request.setAddress(data.address);

        // ❌ REMOVE token here (no auth needed)
        client.register(request, {}, (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    });
};

export const loginUser = async (data: { email: string; password: string }) => {
    return new Promise((resolve, reject) => {
        const request = new LoginRequest();
        request.setEmail(data.email);
        request.setPassword(data.password);

        // ❌ REMOVE token here (no auth needed)
        client.login(request, {}, (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    });
};
