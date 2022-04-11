describe('API Test',() => {

    Cypress.config('baseUrl', 'http://localhost:8081/')

    it("It should be Success Login",()=>{
        cy.request('POST','http://localhost:8080/login').then((response)=>{
            expect(response).to.have.property('status',200)
        })
    })

    it("It should be Get All Materials",() =>{
        cy.request('GET','/material_list').then((response)=>{
            expect(response).to.have.property('status',200)
        
        })
    })

    it("It should be Get All Menus",() =>{
        cy.request('GET','/menu_list_owner').then((response)=>{
            expect(response).to.have.property('status',200)
        
        })
    })

    it("It should be Get All Recipes",() =>{
        cy.request('GET','/recipe_list').then((response)=>{
            expect(response).to.have.property('status',200)
        
        })
    })

    it("It should be Get Specific Menus for Owner",() =>{
        cy.request('GET','/menu_list_search/menu/Ayam').then((response)=>{
            expect(response).to.have.property('status',200)
        
        })
    })

    it("It should be Get Specific Materials",() =>{
        cy.request('GET','/material_list_search/material/Cabai%20Merah').then((response)=>{
            expect(response).to.have.property('status',200)
        
        })
    })

    it("it should be success adding menu",()=>{
        cy.request('POST','localhost:8080/OwnerManager/insertmenu').then((response)=>{
            expect(response).to.have.property('status',200)
        })
    })
   

    it("It should be Get Specific menu description by Id",()=>{
        cy.request('GET','/menu_recipe/29').then((response)=>{
            expect(response).to.have.property('status',200)
        })
    })

   
    it("It should be Get All Menus for Chef",() =>{
        cy.request('GET','/menu_list_chef').then((response)=>{
            expect(response).to.have.property('status',200)
        
        })
    })

})