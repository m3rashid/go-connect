import { AxiosRequestHeaders } from "axios";

const token = localStorage.getItem("connect-token");
export const headers: AxiosRequestHeaders = {
  "Content-type": "application/json",
  authorization: token || false,
};
