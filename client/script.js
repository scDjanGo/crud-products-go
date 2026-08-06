const BASE_URL = "http://127.0.0.1:8080/api/v1"


async function Start() {
    const PRODUCTS = await FetchListData()
    const container = CreateContainer()



    PRODUCTS.forEach(item => {
        container.append(CreateElem(item))
    })

    CreateProductForm()



}
Start()




async function FetchListData() {
    try {
        const response = await fetch(`${BASE_URL}/products`)
        const data = await response.json()
        return data
    } catch (err) {
        console.error(err);
    }
}




function CreateContainer() {
    const container = document.createElement("div")

    container.classList.add("container")

    document.body.prepend(container)

    return container


}



function CreateElem(productData) {
    const element = document.createElement("div")
    element.setAttribute("data-id", productData.id)


    const title = document.createElement("h2")
    title.textContent = productData.title


    const description = document.createElement("p")
    description.textContent = productData.description


    element.appendChild(title)
    element.appendChild(description)
    return element

}







function CreateProductForm() {
    const formElem = document.createElement("form")
    formElem.classList.add("create-form")

    const fileds = [
        {
            name: "title",
            type: "text",
            placeholder: "Write title"
        },
        {
            name: "description",
            type: "text",
            placeholder: "Write description"
        },
        {
            name: "price",
            type: "number",
            placeholder: "Write price"
        },
        {
            name: "discount",
            type: "number",
            placeholder: "Write discount"
        },
        {
            name: "quantity_in_stock",
            type: "number",
            placeholder: "Write quantity in stock"
        },
    ]


    fileds.forEach(item => {
        const label = document.createElement("label")
        label.setAttribute("for", item.name)

        const tag = document.createElement("p")
        label.appendChild(tag)
        tag.textContent = item.name


        const input = document.createElement("input")
        for (let key in item) {
            input.setAttribute(key, item[key])
        }

        label.appendChild(input)
        formElem.appendChild(label)
    })


    const button = document.createElement("button")
    button.setAttribute("type", "submit")
    button.textContent = "create product"

    formElem.append(button)
    formElem.addEventListener("submit", createProduct)

    document.body.prepend(formElem)
}



async function createProduct(event) {
    event.preventDefault()

    const formData = Object.fromEntries(new FormData(event.currentTarget).entries())


    try {
        const response = await fetch(`${BASE_URL}/products`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })

        const data = await response.json()
        const container = document.querySelector(".container")
        container.prepend(CreateElem(data))

    } catch (err) {

    }
}