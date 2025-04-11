import { userClient } from './Grpc.ts';
import {
    RegisterRequest,
    //RegisterResponse,
    LoginRequest,
    // LoginResponse,
} from '../proto/user';

export interface RegisterUserData {
    full_name: string;
    email: string;
    password: string;
    phone: string;
    date_of_birth: string;
    address: string;
}

export interface LoginUserData {
    email: string;
    password: string;
}

export interface RegisterResponseData {
    message: string;
    status: number;
}

export interface LoginResponseData {
    token: string;
    status: number;
}

export const registerUser = async (
    data: RegisterUserData
): Promise<RegisterResponseData> => {
    const request: RegisterRequest = {
        fullName: data.full_name,
        email: data.email,
        password: data.password,
        phone: data.phone,
        dateOfBirth: data.date_of_birth,
        address: data.address,
    };

    const response = await userClient.Register(request);

    return {
        message: response.message,
        status: response.status,
    };
};

export const loginUser = async (
    data: LoginUserData
): Promise<LoginResponseData> => {
    const request: LoginRequest = {
        email: data.email,
        password: data.password,
    };

    const response = await userClient.Login(request);

    return {
        token: response.token,
        status: response.status,
    };
};
