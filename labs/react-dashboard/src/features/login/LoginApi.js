import axiosInstance from "../../utils/axiosInstance";

const LoginApi = async (payload) => {
  try {
    const response = await axiosInstance.post("/login", payload, {
      withAuth: false,
    });

    console.debug("Login Response: ", response);

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    console.error("Login error: ", error);

    const message =
      error.response?.data?.error || "Login failed. Please try again.";

    return {
      success: false,
      error: message,
    };
  }
};

export default LoginApi;
