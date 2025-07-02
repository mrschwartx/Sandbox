import { Stack, Typography, useMediaQuery } from "@mui/material";
import BannerLogin from "../../assets/images/illustrations/login-illustration.png";
import RegisterForm from "./RegisterForm";

const Register = () => {
  const tablet = useMediaQuery("(max-width: 920px)");

  return (
    <Stack
      direction={"row"}
      sx={{
        height: "100vh",
      }}
    >
      {!tablet && (
        <Stack
          spacing={12}
          sx={{
            width: "40vw",
            justifyContent: "center",
            alignItems: "center",
            background:
              "linear-gradient(0deg, rgba(0, 109, 57, 0.05), rgba(0, 109, 57, 0.05)), #FBFDF7",
            px: 12,
          }}
        >
          <img src={BannerLogin} alt="banner-login" width={"100%"} />
          <Stack>
            <Typography textAlign={"center"} variant="h4">
              Example Dashboard
            </Typography>
          </Stack>
        </Stack>
      )}
      <RegisterForm />
    </Stack>
  );
};

export default Register;
