import { useEffect } from "react";
import { useState } from "react";
import { Navigate, Outlet, useLocation } from "react-router";
import { Helmet } from "react-helmet";
import AuthService from "../services/AuthService";
import { getHeading } from "../utils/getHeading";

const TITLE = "Dashboard";

// User cannot access, if user is not login
const PrivateRoute = () => {
  const [title, setTitle] = useState(TITLE);
  const { pathname } = useLocation();

  useEffect(() => {
    const heading = getHeading(pathname);
    if (heading === "") {
      setTitle(TITLE);
    } else {
      setTitle(heading);
    }
  }, [pathname]);

  if (!AuthService.isAuthorized()) {
    return <Navigate to="/login" replace />;
  }

  return (
    <>
      <Helmet>
        <meta charSet="utf-8" />
        <title>{title}</title>
      </Helmet>
      <Outlet />
    </>
  );
};

export default PrivateRoute;
