import SearchProduct from "@/components/SearchProduct/SearchProduct"

async function Items(params: {
    searchParams: { [key: string]: string | string[] | undefined }
}) {
    const searchQuery = typeof params !== 'undefined' && typeof params.searchParams.search === 'string' ? params.searchParams.search : ''
    const items = await fetch('http://localhost:8080/api/items?search=' + encodeURIComponent(searchQuery)).then((res) => res.json())
    return <>
        {/* <pre>{JSON.stringify(items, null, 4)}</pre> */}
        <SearchProduct product={items.items.map((i: any) => ({
            id: i.id,
            title: i.title,
            price: {
                currency: i.price.currency,
                amount: i.price.amount,
                decimals: i.price.decimals
            },
            picture: i.picture
        }))} />

    </>
}

export default Items
