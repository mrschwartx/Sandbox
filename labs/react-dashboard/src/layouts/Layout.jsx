import { Helmet } from "react-helmet";

const Layout = ({ title = "Dashboard", children }) => {
  return (
    <>
      <Helmet>
        <meta charSet="utf-8" />
        <title>{title}</title>
      </Helmet>

      {children}
    </>
  );
};

export default Layout;
