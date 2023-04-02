import { render, screen } from '@testing-library/react'
import Alert from './alert'
import '@testing-library/jest-dom'

/**
 * @jest-environment jsdom
 */

describe('Alert', () => {
  it('renders particular text when authenticated', () => {
    render(<Alert isAuthenticated={true} />)
    expect(screen.getByText('You are currently authenticated')).toBeVisible()
  })
})