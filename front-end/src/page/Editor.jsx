import React from 'react';
import clsx from 'clsx';
import {
  $getRoot,
  $getSelection,
  $isRangeSelection,
  FORMAT_TEXT_COMMAND,
  FORMAT_ELEMENT_COMMAND,
  UNDO_COMMAND,
  REDO_COMMAND,
} from 'lexical';
import { LexicalComposer } from '@lexical/react/LexicalComposer';
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import { RichTextPlugin } from '@lexical/react/LexicalRichTextPlugin';
import { ContentEditable } from '@lexical/react/LexicalContentEditable';
import { OnChangePlugin } from '@lexical/react/LexicalOnChangePlugin';
import { HistoryPlugin } from '@lexical/react/LexicalHistoryPlugin';
import { mergeRegister } from '@lexical/utils';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBold, faItalic, faStrikethrough, faUnderline, faAlignLeft, faAlignRight, faAlignCenter, faAlignJustify, faRotateLeft, faRotateRight, faRotate } from '@fortawesome/free-solid-svg-icons';
import "./Editor.css";


export const Editor = (props) => {

  function onChange(state) {
    props.onChange()
    state.read(() => {
      const root = $getRoot();
      const selection = $getSelection();
  
      console.log(selection);
    });
  }
  

  // about editor: I do a quick copy
  // I will review it when I have time
  const initialConfig = {};
  return (
    <div className="bg-white relative rounded-sm shadow-sm border border-gray-200">
      
      <LexicalComposer
        initialConfig={{
          theme: {
            paragraph: 'mb-1',
            rtl: 'text-right',
            ltr: 'text-left',
            text: {
              bold: 'font-bold',
              italic: 'italic',
              underline: 'underline',
              strikethrough: 'line-through',
            },
          },
          onError(error) {
            throw error;
          },
        }}
      >
        
        <RichTextPlugin
          contentEditable={
            <ContentEditable id="editor-content" className="outline-none py-[15px] px-2.5 resize-none text-ellipsis"/>
          }
        />
        <Toolbar />
        <OnChangePlugin onChange={onChange} />
        <HistoryPlugin />
      </LexicalComposer>
    </div>
  );
};

const Toolbar = () => {
  const [editor] = useLexicalComposerContext();
  const [isBold, setIsBold] = React.useState(false);
  const [isItalic, setIsItalic] = React.useState(false);
  const [isStrikethrough, setIsStrikethrough] = React.useState(false);
  const [isUnderline, setIsUnderline] = React.useState(false);

  const updateToolbar = React.useCallback(() => {
    const selection = $getSelection();

    if ($isRangeSelection(selection)) {
      setIsBold(selection.hasFormat('bold'));
      setIsItalic(selection.hasFormat('italic'));
      setIsStrikethrough(selection.hasFormat('strikethrough'));
      setIsUnderline(selection.hasFormat('underline'));
    }
  }, [editor]);

  React.useEffect(() => {
    return mergeRegister(
      editor.registerUpdateListener(({ editorState }) => {
        editorState.read(() => {
          updateToolbar();
        });
      })
    );
  }, [updateToolbar, editor]);

  return (
    <div className="absolute z-20 shadow bottom--10 left-1/2 transform -translate-x-1/2 min-w-52 h-10 px-2 py-2 bg-[#1b2733] mb-4 space-x-2 flex items-center">
      <button
        className={clsx(
          'px-1 hover:bg-gray-700 transition-colors duration-100 ease-in',
          isBold ? 'bg-gray-700' : 'bg-transparent'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'bold');
        }}
      >
        <FontAwesomeIcon
          icon={faBold}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 hover:bg-gray-700 transition-colors duration-100 ease-in',
          isStrikethrough ? 'bg-gray-700' : 'bg-transparent'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'strikethrough');
        }}
      >
        <FontAwesomeIcon
          icon={faStrikethrough}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 hover:bg-gray-700 transition-colors duration-100 ease-in',
          isItalic ? 'bg-gray-700' : 'bg-transparent'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'italic');
        }}
      >
        <FontAwesomeIcon
          icon={faItalic}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 hover:bg-gray-700 transition-colors duration-100 ease-in',
          isUnderline ? 'bg-gray-700' : 'bg-transparent'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_TEXT_COMMAND, 'underline');
        }}
      >
        <FontAwesomeIcon
          icon={faUnderline}
          className="text-white w-3.5 h-3.5"
        />
      </button>

      <span className="w-[1px] bg-gray-600 block h-full"></span>

      <button
        className={clsx(
          'px-1 bg-transparent hover:bg-gray-700 transition-colors duration-100 ease-in'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'left');
        }}
      >
        <FontAwesomeIcon
          icon={faAlignLeft}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 bg-transparent hover:bg-gray-700 transition-colors duration-100 ease-in'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'center');
        }}
      >
        <FontAwesomeIcon
          icon={faAlignCenter}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 bg-transparent hover:bg-gray-700 transition-colors duration-100 ease-in'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'right');
        }}
      >
        <FontAwesomeIcon
          icon={faAlignRight}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 bg-transparent hover:bg-gray-700 transition-colors duration-100 ease-in'
        )}
        onClick={() => {
          editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, 'justify');
        }}
      >
        <FontAwesomeIcon
          icon={faAlignJustify}
          className="text-white w-3.5 h-3.5"
        />
      </button>

      <span className="w-[1px] bg-gray-600 block h-full"></span>

      <button
        className={clsx(
          'px-1 bg-transparent hover:bg-gray-700 transition-colors duration-100 ease-in'
        )}
        onClick={() => {
          editor.dispatchCommand(UNDO_COMMAND);
        }}
      >
        <FontAwesomeIcon
          icon={faRotateLeft}
          className="text-white w-3.5 h-3.5"
        />
      </button>
      <button
        className={clsx(
          'px-1 bg-transparent hover:bg-gray-700 transition-colors duration-100 ease-in'
        )}
        onClick={() => {
          editor.dispatchCommand(REDO_COMMAND);
        }}
      >
        <FontAwesomeIcon
          icon={faRotateRight}
          className="text-white w-3.5 h-3.5"
        />
      </button>
    </div>
  );
};
