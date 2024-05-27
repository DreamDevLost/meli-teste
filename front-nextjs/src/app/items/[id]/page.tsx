import Breadcrumb from "@/components/Breadcrumb/Breadcrumb";
import Navbar from "@/components/Navbar/Navbar";
import Product from "@/components/Product/Product";

export default async function Page({ params }: { params: { id: string } }) {
    const { item } = await fetch(`http://localhost:8080/api/items/${params.id}`).then((res) => res.json())

    return <>
        <Breadcrumb links={[
            { name: 'Eletronicos', url: 'http://localhost:8080' },
            { name: 'Iphone', url: 'http://localhost:8080' },
            { name: '256GB' },
        ]} />
        <Product item={{
            id: item.id,
            title: item.title,
            price: item.price,
            picture: item.picture,
            condition: item.condition,
            free_shipping: item.free_shipping,
            sold_quantity: item.sold_quantity,
            description: item.description
        }} />
    </>
}