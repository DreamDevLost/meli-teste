import Image from "next/image";
import styles from "./page.module.scss";
import SearchProduct from "@/components/SearchProduct/SearchProduct";
import Breadcrumb from "@/components/Breadcrumb/Breadcrumb";
import Product from "@/components/Product/Product";

export default function Home() {
  return (<>
    <Breadcrumb links={[
      { name: 'Eletronicos', url: 'http://localhost:8080' },
      { name: 'Iphone', url: 'http://localhost:8080' },
      { name: '256GB' },
    ]} />
  </>
  );
}
