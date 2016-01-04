package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSignalTransition struct {
	QAbstractTransition
}

type QSignalTransition_ITF interface {
	QAbstractTransition_ITF
	QSignalTransition_PTR() *QSignalTransition
}

func PointerFromQSignalTransition(ptr QSignalTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalTransition_PTR().Pointer()
	}
	return nil
}

func NewQSignalTransitionFromPointer(ptr unsafe.Pointer) *QSignalTransition {
	var n = new(QSignalTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSignalTransition_") {
		n.SetObjectName("QSignalTransition_" + qt.Identifier())
	}
	return n
}

func (ptr *QSignalTransition) QSignalTransition_PTR() *QSignalTransition {
	return ptr
}

func NewQSignalTransition(sourceState QState_ITF) *QSignalTransition {
	defer qt.Recovering("QSignalTransition::QSignalTransition")

	return NewQSignalTransitionFromPointer(C.QSignalTransition_NewQSignalTransition(PointerFromQState(sourceState)))
}

func NewQSignalTransition2(sender QObject_ITF, signal string, sourceState QState_ITF) *QSignalTransition {
	defer qt.Recovering("QSignalTransition::QSignalTransition")

	return NewQSignalTransitionFromPointer(C.QSignalTransition_NewQSignalTransition2(PointerFromQObject(sender), C.CString(signal), PointerFromQState(sourceState)))
}

func (ptr *QSignalTransition) Event(e QEvent_ITF) bool {
	defer qt.Recovering("QSignalTransition::event")

	if ptr.Pointer() != nil {
		return C.QSignalTransition_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSignalTransition) EventTest(event QEvent_ITF) bool {
	defer qt.Recovering("QSignalTransition::eventTest")

	if ptr.Pointer() != nil {
		return C.QSignalTransition_EventTest(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSignalTransition) ConnectOnTransition(f func(event *QEvent)) {
	defer qt.Recovering("connect QSignalTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onTransition", f)
	}
}

func (ptr *QSignalTransition) DisconnectOnTransition() {
	defer qt.Recovering("disconnect QSignalTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onTransition")
	}
}

//export callbackQSignalTransitionOnTransition
func callbackQSignalTransitionOnTransition(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalTransition::onTransition")

	if signal := qt.GetSignal(C.GoString(ptrName), "onTransition"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	}

}

func (ptr *QSignalTransition) OnTransition(event QEvent_ITF) {
	defer qt.Recovering("QSignalTransition::onTransition")

	if ptr.Pointer() != nil {
		C.QSignalTransition_OnTransition(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSignalTransition) OnTransitionDefault(event QEvent_ITF) {
	defer qt.Recovering("QSignalTransition::onTransition")

	if ptr.Pointer() != nil {
		C.QSignalTransition_OnTransitionDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSignalTransition) SenderObject() *QObject {
	defer qt.Recovering("QSignalTransition::senderObject")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalTransition_SenderObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSignalTransition) ConnectSenderObjectChanged(f func()) {
	defer qt.Recovering("connect QSignalTransition::senderObjectChanged")

	if ptr.Pointer() != nil {
		C.QSignalTransition_ConnectSenderObjectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "senderObjectChanged", f)
	}
}

func (ptr *QSignalTransition) DisconnectSenderObjectChanged() {
	defer qt.Recovering("disconnect QSignalTransition::senderObjectChanged")

	if ptr.Pointer() != nil {
		C.QSignalTransition_DisconnectSenderObjectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "senderObjectChanged")
	}
}

//export callbackQSignalTransitionSenderObjectChanged
func callbackQSignalTransitionSenderObjectChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSignalTransition::senderObjectChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "senderObjectChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSignalTransition) SetSenderObject(sender QObject_ITF) {
	defer qt.Recovering("QSignalTransition::setSenderObject")

	if ptr.Pointer() != nil {
		C.QSignalTransition_SetSenderObject(ptr.Pointer(), PointerFromQObject(sender))
	}
}

func (ptr *QSignalTransition) SetSignal(signal QByteArray_ITF) {
	defer qt.Recovering("QSignalTransition::setSignal")

	if ptr.Pointer() != nil {
		C.QSignalTransition_SetSignal(ptr.Pointer(), PointerFromQByteArray(signal))
	}
}

func (ptr *QSignalTransition) Signal() *QByteArray {
	defer qt.Recovering("QSignalTransition::signal")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QSignalTransition_Signal(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSignalTransition) ConnectSignalChanged(f func()) {
	defer qt.Recovering("connect QSignalTransition::signalChanged")

	if ptr.Pointer() != nil {
		C.QSignalTransition_ConnectSignalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalChanged", f)
	}
}

func (ptr *QSignalTransition) DisconnectSignalChanged() {
	defer qt.Recovering("disconnect QSignalTransition::signalChanged")

	if ptr.Pointer() != nil {
		C.QSignalTransition_DisconnectSignalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalChanged")
	}
}

//export callbackQSignalTransitionSignalChanged
func callbackQSignalTransitionSignalChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSignalTransition::signalChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "signalChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSignalTransition) DestroyQSignalTransition() {
	defer qt.Recovering("QSignalTransition::~QSignalTransition")

	if ptr.Pointer() != nil {
		C.QSignalTransition_DestroyQSignalTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSignalTransition) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSignalTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSignalTransition) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSignalTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSignalTransitionTimerEvent
func callbackQSignalTransitionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalTransition::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQSignalTransitionFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSignalTransition) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QSignalTransition::timerEvent")

	if ptr.Pointer() != nil {
		C.QSignalTransition_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSignalTransition) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QSignalTransition::timerEvent")

	if ptr.Pointer() != nil {
		C.QSignalTransition_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSignalTransition) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSignalTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSignalTransition) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSignalTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSignalTransitionChildEvent
func callbackQSignalTransitionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalTransition::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQSignalTransitionFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QSignalTransition) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QSignalTransition::childEvent")

	if ptr.Pointer() != nil {
		C.QSignalTransition_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSignalTransition) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QSignalTransition::childEvent")

	if ptr.Pointer() != nil {
		C.QSignalTransition_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSignalTransition) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSignalTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSignalTransition) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSignalTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSignalTransitionCustomEvent
func callbackQSignalTransitionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalTransition::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQSignalTransitionFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QSignalTransition) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QSignalTransition::customEvent")

	if ptr.Pointer() != nil {
		C.QSignalTransition_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSignalTransition) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QSignalTransition::customEvent")

	if ptr.Pointer() != nil {
		C.QSignalTransition_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
