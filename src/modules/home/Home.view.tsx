import HangBoard from "./assets/hangboard.jpg";
import WoodShop from "./assets/wood_shop.jpg";
import Printing from "./assets/printing.jpg";
import Logo from "@assets/logo.png";
import { ServiceCard, type ServiceCardProps } from "./components";
import cx from "classnames";
import { PAGES } from "src/lib/routes";

export default () => {
  const navLinks: ServiceCardProps[] = [
    { text: "Climbing", href: PAGES.climbing, src: HangBoard },
    { text: "Woodshop", href: PAGES.woodShop, src: WoodShop },
  ];

  const printingService: ServiceCardProps = {
    text: "3D Printings",
    href: PAGES.printing,
    src: Printing,
  };

  return (
    <div className="h-full flex flex-col justify-center items-center gap-10">
      <img
        src={Logo}
        alt="Johnson and Brooks"
        className="w-36 md:w-72 h-auto mx-auto"
      />

      <section className="flex flex-col items-center justify-center gap-3 mb-10">
        <h2 className="text-3xl font-semibold fg-accent">Wood Work</h2>

        <ul className={cx("flex justify-center flex-wrap items-center gap-5")}>
          {navLinks.map((props) => (
            <li key={props.href} className="w-60 md:w-72 lg:w-96">
              <ServiceCard {...props} />
            </li>
          ))}
        </ul>
      </section>

      <section className="flex flex-col items-center justify-center gap-3">
        <h2 className="text-3xl font-semibold fg-accent">3D Printings</h2>
        <div className="w-1/2 max-w-6xl">
          <ServiceCard {...printingService} />
        </div>
      </section>
    </div>
  );
};
