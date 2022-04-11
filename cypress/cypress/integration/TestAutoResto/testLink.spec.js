describe('Link Page Test',()=>{
Cypress.config('baseUrl', 'http://localhost:8081/')

    it.only('The Home page links should be work',()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/')
        cy.url().should('include','/')
      
    })

    it.only('The Login page links should work',()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/login')
        cy.contains('Login').click()
        cy.url().should('include','/login')
        cy.contains('Login').should('exist')
    })

    it.only('The Homepage should be accessible ',()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/home')
        cy.url().should('include','/home')
    })
    //Owner Screen
    it.only("The Add Menu Screen should be accessible",()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/add_menu')
        cy.contains('Insert').click()
        cy.url().should('include','/add_menu')
        cy.contains('Insert').should('exist')
    })

    it.only("The Search Menu Screen should be accessible",()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/search_menu')
        cy.url().should('include','/search_menu')
        cy.contains('Search').should('exist')
    })

    it.only("The Search Material screen should be accessible",()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/search_material')
        cy.url().should('include','/search_material')
        cy.contains('Search').should('exist')  
    })

    it.only("The material list screen should be accessible",()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/material_list')
        cy.url().should('include','/material_list')
        cy.contains('Edit').should('exist')
        cy.contains('Delete').should('exist')  
    })

    it.only("The menu list screen should be accessible",()=>{
        cy.viewport(1280,720)
        cy.visit('http://localhost:8081/menu_list_owner')
        cy.url().should('include','/menu_list')
        cy.contains('Update').should('exist')
        cy.contains('Delete').should('exist')
    })

    it.only("The Recipe list screen shoult be accessible",()=>{
        cy.viewport(1280,270)
        cy.visit('http://localhost:8081/recipe_list')
        cy.url().should('include','/recipe_list')
    })

    it.only("The Update specific material list screen should be accesisble",()=>{
        cy.viewport(1280,900)
        cy.visit('http://localhost:8081/update_material/97')
        cy.url().should('include','/update_material/97')
        cy.contains('Edit').should('exist')

    })

    //Chef Screen
    it.only("The add recipe detail screen should be accessible",()=>{
        cy.viewport(1280,900)
        cy.visit('http://localhost:8081/add_detail_recipe')
        cy.url().should('include','/add_detail_recipe')
        cy.contains('Insert').should('exist')
    })

    it.only("The form menu recipe should be accessible",()=>{
        cy.viewport(1280,270)
        cy.visit('http://localhost:8081/form_recipe_menu')
        cy.url().should('include','/form_recipe_menu')
        cy.contains('submit').should('exist')
    })

    it.only("The spesific description menu should be accessible",()=>{
        cy.viewport(1280,400)
        cy.visit('http://localhost:8081/menu_recipe/29')
        cy.url().should('include','/menu_recipe/29')

    })

    //Inventory Staff Screen
    it.only("The add material screen should be accessible",()=>{
        cy.visit('http://localhost:8081/add_material')
        cy.url().should('include','/add_material')
        cy.contains('Insert').should('exist')
        cy.contains('Insert Material').should('exist')
    })
   
})



