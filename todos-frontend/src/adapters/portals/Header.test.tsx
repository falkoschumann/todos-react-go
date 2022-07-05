import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

import Header from './Header';

const handleAddTodo = jest.fn();

beforeEach(() => {
  handleAddTodo.mockReset();
});

describe('Header', () => {
  it('sends title for new todo.', async () => {
    userEvent.setup();
    render(<Header onAddTodo={handleAddTodo} />);
    const newTodoElement: HTMLInputElement = screen.getByPlaceholderText('What needs to be done?');

    await userEvent.type(newTodoElement, 'Taste JavaScript');
    await userEvent.keyboard('[Enter]');

    expect(newTodoElement.value).toEqual('');
    expect(handleAddTodo).toBeCalledWith('Taste JavaScript');
  });

  it('sends trimed title.', async () => {
    userEvent.setup();
    render(<Header onAddTodo={handleAddTodo} />);
    const newTodoElement = screen.getByPlaceholderText('What needs to be done?');

    await userEvent.type(newTodoElement, '  Taste JavaScript ');
    await userEvent.keyboard('[Enter]');

    expect(handleAddTodo).toBeCalledWith('Taste JavaScript');
  });
});
