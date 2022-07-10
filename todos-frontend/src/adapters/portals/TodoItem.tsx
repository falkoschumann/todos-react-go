import { ChangeEvent, KeyboardEvent, useEffect, useRef, useState } from 'react';
import classNames from 'classnames';

import { Todo } from '../../domain/todo';

import { usePrevious } from './hooks';

type TodoItemProps = Readonly<{
  editing?: boolean;
  todo: Todo;
  onCancel?: () => void;
  onDestroy?: () => void;
  onEdit?: () => void;
  onSave?: (newTitle: string) => void;
  onToggle?: () => void;
}>;

function TodoItem({
  editing = false,
  todo,
  onCancel,
  onDestroy,
  onEdit,
  onSave,
  onToggle,
}: TodoItemProps) {
  const [editText, setEditText] = useState(todo.title);
  const editorRef = useRef<HTMLInputElement>(null);
  const prevEditing = usePrevious(editing);

  function handleEdit() {
    onEdit?.();
    setEditText(todo.title);
  }

  function handleSubmit() {
    const title = editText.trim();
    onSave?.(title);
    setEditText(title);
  }

  function handleKeyDown(event: KeyboardEvent<HTMLInputElement>) {
    switch (event.key) {
      case 'Enter':
        handleSubmit();
        break;
      case 'Escape':
        setEditText(todo.title);
        onCancel?.();
        break;
      default:
        break;
    }
  }

  function handleChange(event: ChangeEvent<HTMLInputElement>) {
    setEditText(event.target.value);
  }

  useEffect(() => {
    if (prevEditing || !editing || editorRef.current == null) {
      return;
    }

    editorRef.current.focus();
    editorRef.current.setSelectionRange(0, editorRef.current.value.length);
  }, [editing, prevEditing]);

  return editing ? (
    <li className="relative text-2xl border-b border-solid border-gray-200 last:border-b-0">
      <input
        ref={editorRef}
        className="w-full p-4 pl-16 font-light text-2xl shadow-inner focus:outline-red-700/40"
        value={editText}
        onBlur={handleSubmit}
        onChange={handleChange}
        onKeyDown={handleKeyDown}
      />
    </li>
  ) : (
    <li className="group relative text-2xl border-b border-solid border-gray-200 last:border-b-0">
      <input
        data-testid="completed"
        type="checkbox"
        checked={todo.completed}
        onChange={onToggle}
        className="absolute top-0 bottom-0 left-3 w-6 h-6 my-auto"
      />
      <label
        onDoubleClick={handleEdit}
        className={classNames({
          'block p-4 pl-16 font-normal break-all': true,
          'text-gray-400 line-through': todo.completed,
          'text-gray-600': !todo.completed,
        })}
      >
        {todo.title}
      </label>
      <button
        className="hidden group-hover:block absolute top-0 right-5 bottom-0 w-10 h-10 my-auto text-gray-400 hover:text-red-300"
        onClick={onDestroy}
      >
        x
      </button>
    </li>
  );
}

export default TodoItem;
