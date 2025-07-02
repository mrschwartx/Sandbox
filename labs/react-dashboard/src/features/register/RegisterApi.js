import axiosInstance from "../../utils/axiosInstance";

const RegisterApi = async (payload) => {
  try {
    const response = await axiosInstance.post("/register", payload, {
      withAuth: false,
    });

    console.debug("Register response:", response);

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    console.error("Register error:", error);

    const message =
      error.response?.data?.error || "Registration failed. Please try again.";

    return {
      success: false,
      error: message,
    };
  }
};

export default RegisterApi;
