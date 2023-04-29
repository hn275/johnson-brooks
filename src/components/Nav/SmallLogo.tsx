import Logo from "@assets/logo_small.svg";

interface Props {
  className?: string;
}

export function SmallLogo(props: Props) {
  return (
    <a href="/" {...props}>
      <img src={Logo} alt="home page" />
    </a>
  );
}
