export const getHeading = (pathname) => {
  const location = pathname.split("/");
  if (location.length > 2) {
    const [, name] = location[1].split("-");
    const heading = `detail ${name}`.toUpperCase();
    return heading;
  }
  return location[1].replaceAll("-", " ").toUpperCase();
};
