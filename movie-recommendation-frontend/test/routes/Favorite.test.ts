import { test, suite, expect } from 'vitest'
import Favorites from '../../src/routes/Favorites.svelte'
import { tick } from 'svelte'
import sinon from 'sinon'





suite('Favorites Component', () => {
    test('should render the component', async () => {
        const host = document.createElement('div')
        document.body.appendChild(host)
        const instance = new Favorites({ target: host })
        expect(instance).toBeTruthy()
        await tick()
        expect(host.innerHTML).toContain('Favorites')
    })

    test('should fetch favorites if jwt token in localStorage', async () => {
        const fetchStub = sinon.stub(global, 'fetch')
        fetchStub.resolves({
            json: () => Promise.resolve({ results: [], total_pages: 0 }),
        })
        localStorage.setItem('jwt', 'testToken')
        const host = document.createElement('div')
        document.body.appendChild(host)
        await new Favorites({ target: host })
        sinon.assert.calledTwice(fetchStub)
        fetchStub.restore()
        localStorage.removeItem('jwt')
    })
})