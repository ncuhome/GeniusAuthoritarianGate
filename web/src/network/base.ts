import axios, { AxiosError } from "axios";

export const BaseURL = `/login/api/`;

export const ErrNetwork = "网络异常";

const api = axios.create({
  baseURL: BaseURL,
});

export function apiV1ErrorHandler(err: AxiosError<any>): any {
  switch (true) {
    case err.name === "CanceledError":
      break;
    case !err || !err.response || !err.response.data:
      err.message = ErrNetwork;
      break;
    default:
      err.message = err.response?.data?.msg;
  }
  return err;
}

api.interceptors.response.use(undefined, (err: AxiosError<any>) => {
  return Promise.reject(apiV1ErrorHandler(err));
});

export { api };
