// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// #include <form.h>
// #include <menu.h>
import "C"

// Form Driver Requests
type FormDriverReq C.int

const (
	REQ_NEXT_PAGE    FormDriverReq = C.REQ_NEXT_PAGE    // next page
	REQ_PREV_PAGE                  = C.REQ_PREV_PAGE    // previous page
	REQ_FIRST_PAGE                 = C.REQ_FIRST_PAGE   // first page
	REQ_LAST_PAGE                  = C.REQ_LAST_PAGE    // last page
	REQ_NEXT_FIELD                 = C.REQ_NEXT_FIELD   // next field
	REQ_PREV_FIELD                 = C.REQ_PREV_FIELD   // previous field
	REQ_FIRST_FIELD                = C.REQ_FIRST_FIELD  // first field
	REQ_LAST_FIELD                 = C.REQ_LAST_FIELD   // last field
	REQ_SNEXT_FIELD                = C.REQ_SNEXT_FIELD  // sorted next field
	REQ_SPREV_FIELD                = C.REQ_SPREV_FIELD  // sorted previous field
	REQ_SFIRST_FIELD               = C.REQ_SFIRST_FIELD // sorted first field
	REQ_SLAST_FIELD                = C.REQ_SLAST_FIELD  // sorted last field
	REQ_LEFT_FIELD                 = C.REQ_LEFT_FIELD   // left field
	REQ_RIGHT_FIELD                = C.REQ_RIGHT_FIELD  // right field
	REQ_UP_FIELD                   = C.REQ_UP_FIELD     // up to a field
	REQ_DOWN_FIELD                 = C.REQ_DOWN_FIELD   // down to a field
	REQ_NEXT_CHAR                  = C.REQ_NEXT_CHAR    // next character in field
	REQ_PREV_CHAR                  = C.REQ_PREV_CHAR    // previous character in field
	REQ_NEXT_LINE                  = C.REQ_NEXT_LINE    // next line
	REQ_PREV_LINE                  = C.REQ_PREV_LINE    // previous line
	REQ_NEXT_WORD                  = C.REQ_NEXT_WORD    // next word
	REQ_PREV_WORD                  = C.REQ_PREV_WORD    // previous word
	REQ_BEG_FIELD                  = C.REQ_BEG_FIELD    // beginning of field
	REQ_END_FIELD                  = C.REQ_END_FIELD    // end of field
	REQ_BEG_LINE                   = C.REQ_BEG_LINE     // beginning of line
	REQ_END_LINE                   = C.REQ_END_LINE     // end of line
	REQ_LEFT_CHAR                  = C.REQ_LEFT_CHAR    // character to the left
	REQ_RIGHT_CHAR                 = C.REQ_RIGHT_CHAR   // character to the right
	REQ_UP_CHAR                    = C.REQ_UP_CHAR      // up a character
	REQ_DOWN_CHAR                  = C.REQ_DOWN_CHAR    // down a character
	REQ_NEW_LINE                   = C.REQ_NEW_LINE     // insert of overlay a new line
	REQ_INS_CHAR                   = C.REQ_INS_CHAR     // insert a blank character at cursor
	REQ_INS_LINE                   = C.REQ_INS_LINE     // insert a blank line at cursor
	REQ_DEL_CHAR                   = C.REQ_DEL_CHAR     // delete character at cursor
	REQ_DEL_PREV                   = C.REQ_DEL_PREV     // delete character before cursor
	REQ_DEL_LINE                   = C.REQ_DEL_LINE     // delete line
	REQ_DEL_WORD                   = C.REQ_DEL_WORD     // delete word
	REQ_CLR_EOL                    = C.REQ_CLR_EOL      // clear from cursor to end of line
	REQ_CLR_EOF                    = C.REQ_CLR_EOF      // clear from cursor to end of field
	REQ_CLR_FIELD                  = C.REQ_CLR_FIELD    // clear field
	REQ_OVL_MODE                   = C.REQ_OVL_MODE     // overlay mode
	REQ_INS_MODE                   = C.REQ_INS_MODE     // insert mode
	REQ_SCR_FLINE                  = C.REQ_SCR_FLINE    // scroll field forward a line
	REQ_SCR_BLINE                  = C.REQ_SCR_BLINE    // scroll field back a line
	REQ_SCR_FPAGE                  = C.REQ_SCR_FPAGE    // scroll field forward a page
	REQ_SCR_BPAGE                  = C.REQ_SCR_BPAGE    // scroll field back a page
	REQ_SCR_FHPAGE                 = C.REQ_SCR_FHPAGE   // scroll field forward half a page
	REQ_SCR_BHPAGE                 = C.REQ_SCR_BHPAGE   // scroll field back half a page
	REQ_SCR_FCHAR                  = C.REQ_SCR_FCHAR    // scroll field forward a character
	REQ_SCR_BCHAR                  = C.REQ_SCR_BCHAR    // scroll field back a character
	REQ_SCR_HFLINE                 = C.REQ_SCR_HFLINE   // horisontal scroll field forward a line
	REQ_SCR_HBLINE                 = C.REQ_SCR_HBLINE   // horisontal scroll field back a line
	REQ_SCR_HFHALF                 = C.REQ_SCR_HFHALF   // horisontal scroll field forward half a line
	REQ_SCR_HBHALF                 = C.REQ_SCR_HBHALF   // horisontal scroll field back half a line
	REQ_VALIDATION                 = C.REQ_VALIDATION   // validate field
	REQ_NEXT_CHOICE                = C.REQ_NEXT_CHOICE  // display next field choice
	REQ_PREV_CHOICE                = C.REQ_PREV_CHOICE  // display previous field choice
)

const (
	FO_VISIBLE  = C.O_VISIBLE  // Field visibility
	FO_ACTIVE   = C.O_ACTIVE   // Field is sensitive/accessable
	FO_PUBLIC   = C.O_PUBLIC   // Typed characters are echoed
	FO_EDIT     = C.O_EDIT     // Editable
	FO_WRAP     = C.O_WRAP     // Line wrapping
	FO_BLANK    = C.O_BLANK    // Clear on entry
	FO_AUTOSKIP = C.O_AUTOSKIP // Skip to next field when current filled
	FO_NULLOK   = C.O_NULLOK   // Blank ok
	FO_STATIC   = C.O_STATIC   // Fixed size
	FO_PASSOK   = C.O_PASSOK   // Field validation
)

// Menu Driver Requests
type MenuDriverReq C.int

const (
	REQ_LEFT          MenuDriverReq = C.REQ_LEFT_ITEM
	REQ_RIGHT                       = C.REQ_RIGHT_ITEM
	REQ_UP                          = C.REQ_UP_ITEM
	REQ_DOWN                        = C.REQ_DOWN_ITEM
	REQ_ULINE                       = C.REQ_SCR_ULINE
	REQ_DLINE                       = C.REQ_SCR_DLINE
	REQ_PAGE_DOWN                   = C.REQ_SCR_DPAGE
	REQ_PAGE_UP                     = C.REQ_SCR_UPAGE
	REQ_FIRST                       = C.REQ_FIRST_ITEM
	REQ_LAST                        = C.REQ_LAST_ITEM
	REQ_NEXT                        = C.REQ_NEXT_ITEM
	REQ_PREV                        = C.REQ_PREV_ITEM
	REQ_TOGGLE                      = C.REQ_TOGGLE_ITEM
	REQ_CLEAR_PATTERN               = C.REQ_CLEAR_PATTERN
	REQ_BACK_PATTERN                = C.REQ_BACK_PATTERN
	REQ_NEXT_MATCH                  = C.REQ_NEXT_MATCH
	REQ_PREV_MATCH                  = C.REQ_PREV_MATCH
)

// Menu Options
const (
	O_ONEVALUE   = C.O_ONEVALUE   // Only one item can be selected
	O_SHOWDESC   = C.O_SHOWDESC   // Display item descriptions
	O_ROWMAJOR   = C.O_ROWMAJOR   // Display in row-major order
	O_IGNORECASE = C.O_IGNORECASE // Ingore case when pattern-matching
	O_SHOWMATCH  = C.O_SHOWMATCH  // Move cursor to item when pattern-matching
	O_NONCYCLIC  = C.O_NONCYCLIC  // Don't wrap next/prev item
)

// Menu Item Options
const O_SELECTABLE = C.O_SELECTABLE
